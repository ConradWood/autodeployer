package main

// instruct the autodeployer on a given server to download & deploy stuff

import (
	"context"
	"errors"
	"flag"
	"fmt"

	"golang.conradwood.net/apis/common"
	pb "golang.conradwood.net/apis/deploymonkey"
	dc "golang.conradwood.net/deploymonkey/common"
	"golang.conradwood.net/deploymonkey/config"
	"golang.conradwood.net/deploymonkey/suggest"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/client"
	"golang.conradwood.net/go-easyops/utils"

	//	"google.golang.org/grpc"
	"os"
	"strings"
	"time"
)

// static variables for flag parser
var (
	continue_on_error = flag.Bool("continue_on_error", false, "if true ignore deploy errors and continue to next deployment")
	del_version       = flag.Uint64("delete_version", 0, "if not 0, undeploy and delete this version")
	depllocal         = flag.Bool("deploy_local", false, "deploy the files in current git repository on local machine (using deploy.yaml in current directory)")
	depllist          = flag.Bool("deployments", false, "list current deployments")
	list_deployers    = flag.Bool("deployers", false, "list known autodeployers")
	dosuggest         = flag.Bool("suggest", false, "suggest fixes")
	applysuggest      = flag.Bool("apply_suggest", false, "suggest & applyfixes")
	short             = flag.Bool("short", false, "short listing")
	p_newbuild        = flag.Bool("newbuild", false, "if true, simulate a new build, as triggered by buildrepo")
	filename          = flag.String("configfile", "", "the yaml config file to submit to server")
	namespace         = flag.String("namespace", "", "namespace of the group to update")
	groupname         = flag.String("groupname", "", "groupname of the group to update")
	repository        = flag.Uint64("repository", 0, "repository of the app in the group to update")
	buildid           = flag.Int("buildid", 0, "the new buildid of the app in the group to update")
	binary            = flag.String("binary", "", "the binary of the app in the group to update")
	apply_version     = flag.Int("apply_version", 0, "(re-)apply a given version (expects `versionid`)")
	applyall          = flag.Bool("apply_all", false, "reapply ALL groups")
	applypending      = flag.Bool("apply_pending", false, "reapply any pending group versions")
	list              = flag.String("list", "", "list this `repository` previous versions")
	deployers         = flag.Bool("list_deployers", false, "if true list all known autodeployers")
	undeploy_app      = flag.Int("undeploy_version", 0, "undeploy applications of a given version (expects `versionid`)")
	print_sample      = flag.Bool("print_sample", false, "print a sample deploy.yaml")
	depl              pb.DeployMonkeyClient
	deploy_timeout    = flag.Duration("deploy_timeout", time.Duration(30)*time.Second, "deployrequests timeout")
)

func main() {
	flag.Parse()

	fmt.Printf("Starting deploymonkey client...\n")
	depl = pb.NewDeployMonkeyClient(client.Connect("deploymonkey.DeployMonkey"))
	if *p_newbuild {
		utils.Bail("failed to do newbuild", newbuild())
		os.Exit(0)
	}
	done := false
	if *del_version != 0 {
		utils.Bail("failed to delete version", delVersion())
		os.Exit(0)
	}
	if *deployers {
		listDeployers()
		os.Exit(0)
	}
	if *print_sample {
		dc.PrintSample()
		os.Exit(0)
	}

	if *depllocal {
		DeployLocal()
		os.Exit(0)
	}

	if *depllist {
		listDeployments()
		os.Exit(0)
	}
	if *dosuggest {
		listSuggestions()
		os.Exit(0)
	}
	if *applysuggest {
		applySuggestions()
		os.Exit(0)
	}

	if *undeploy_app != 0 {
		undeployApplication(*undeploy_app)
		os.Exit(0)
	}
	if *list != "" {
		callListVersions(*list)
		os.Exit(0)
	}

	if *apply_version != 0 {
		applyVersion()
		done = true
	}
	if *namespace != "" {
		if *binary != "" {
			updateApp()
		} else {
			updateRepo()
		}
		done = true
	}
	if !done {
		listConfig()
		fmt.Printf("Nothing to do.\n")
		os.Exit(1)
	}
	os.Exit(0)
}
func bail(err error, msg string) {
	if err == nil {
		return
	}
	fmt.Printf("%s: %s\n", msg, err)
	os.Exit(10)
}

func undeployApplication(ID int) {
	ctx := Context()
	uar := pb.UndeployApplicationRequest{ID: int64(ID)}
	resp, err := depl.UndeployApplication(ctx, &uar)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	if resp.App == nil {
		fmt.Printf("Nothing undeployed! Got the right ID?\n")
	} else {
		fmt.Printf("Undeployed: %d [%s]\n", resp.App.RepositoryID, resp.App.Binary)
	}
	for _, host := range resp.Host {
		fmt.Printf(" on %s\n", host)
	}
}

func callListVersions(repo string) {
	ctx := Context()
	dr := pb.ListVersionByNameRequest{Name: repo}
	resp, err := depl.ListVersionsByName(ctx, &dr)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("%d apps\n", len(resp.Apps))
	t := utils.Table{}
	t.AddHeaders("Version", "Created", "Build", "Binary", "RepoID", "ArtefactID")
	for i, a := range resp.Apps {
		if i > 10 {
			break
		}
		t.AddUint64(uint64(a.VersionID))
		t.AddTimestamp(uint32(a.Created))
		t.AddUint64(a.Application.BuildID)
		t.AddString(a.Application.Binary)
		t.AddUint64(a.Application.RepositoryID)
		t.AddUint64(a.Application.ArtefactID)
		t.NewRow()
		//		fmt.Printf("Version #%d: created %v, Build %d,  binary %s\n", a.VersionID, created, a.Application.BuildID, a.Application.Binary)
	}
	fmt.Println(t.ToPrettyString())
}

func applyVersion() {
	fmt.Printf("Applying Version...\n")
	ctx := Context()

	dr := pb.DeployRequest{VersionID: fmt.Sprintf("%d", *apply_version)}
	resp, err := depl.DeployVersion(ctx, &dr)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Version %d applied (%s)\n", *apply_version, resp)

}

/************************************************
* this is more or less a direct copy of GetConfig
* and should be converted to use GetConfig()
*************************************************/
func listConfig() {
	config, err := config.GetConfig(depl)
	utils.Bail("failed to get config", err)
	t := utils.Table{}
	t.AddHeaders("namespace", "pending#", "deployed#", "AppID", "Build", "RepoID", "ArtefactID", "Binary")
	for _, n := range config.Namespaces() {
		if !matchesArgs(n) {
			continue
		}
		gns := config.Groups(n)
		if (*short) && len(gns) == 1 {
			fmt.Printf("  %s (%d groups) - pending_versionID: #%d\n", n, len(gns), gns[0].PendingVersion)
			continue
		}
		fmt.Printf("  %s (%d groups) pending_version=%d\n", n, len(gns), gns[0].PendingVersion)

		for _, gs := range gns {
			gapps := config.Apps(gs.NameSpace, gs.GroupID)
			marker := ""
			if gs.PendingVersion != gs.DeployedVersion {
				marker = " ** <-- **"
			}
			fmt.Printf("      %s (%d applications)%s\n", gs, len(gapps), marker)
			t.AddString(n)
			t.NewRow()
			for _, app := range gapps {
				t.AddString("")
				t.AddUint64(uint64(gs.PendingVersion))
				t.AddUint64(uint64(gs.DeployedVersion))
				t.AddUint64(app.ID)
				t.AddUint64(app.BuildID)
				t.AddUint64(app.RepositoryID)
				t.AddUint64(app.ArtefactID)
				t.AddString(app.Binary)
				t.NewRow()
				ao := "__"
				if app.AlwaysOn {
					ao = "AO"
				}
				crit := "__"
				if app.Critical {
					crit = "CR"
				}
				fmt.Printf("           [%s%s] %dx App ID=%d, Repo=%d, Binary=%s, BuildID=#%d, %d autoregistrations\n", ao, crit, app.Instances, app.ID, app.RepositoryID, app.Binary, app.BuildID, len(app.AutoRegs))
			}
		}
	}
	fmt.Println(t.ToPrettyString())
}
func matchesArgs(namespace string) bool {
	args := flag.Args()
	if len(args) == 0 {
		return true
	}
	for _, s := range args {
		if strings.Contains(namespace, s) {
			return true
		}
	}
	return false
}

func updateRepo() {
	if *namespace == "" {
		bail(errors.New("Namespace required"), "Cannot update repo")
	}
	if *groupname == "" {
		bail(errors.New("Groupname required"), "Cannot update repo")
	}
	if *repository == 0 {
		bail(errors.New("Repository required"), "Cannot update repo")
	}
	if *buildid == 0 {
		bail(errors.New("BuildID required"), "Cannot update repo")
	}
	fmt.Printf("Updating all apps in repo %d in group %s to buildid %d\n", *repository, *groupname, *buildid)
	ur := pb.UpdateRepoRequest{
		Namespace:    *namespace,
		GroupID:      *groupname,
		RepositoryID: *repository,
		BuildID:      uint64(*buildid),
	}
	ctx := Context()
	resp, err := depl.UpdateRepo(ctx, &ur)
	bail(err, "Failed to update repo")
	fmt.Printf("Response to updaterepo: %v\n", resp)
	return
}

func updateApp() {
	ad := pb.ApplicationDefinition{
		RepositoryID: *repository,
		Binary:       *binary,
		BuildID:      uint64(*buildid),
	}
	uar := pb.UpdateAppRequest{
		GroupID:   *groupname,
		Namespace: *namespace,
		App:       &ad,
	}
	fmt.Printf("Updating app %s\n", *binary)
	ctx := Context()

	resp, err := depl.UpdateApp(ctx, &uar)
	if err != nil {
		fmt.Printf("Failed to update app: %s\n", err)
		return
	}
	fmt.Printf("Response to updateapp: %v\n", resp.Result)
}

func listSuggestions() {
	sr := &pb.SuggestRequest{}
	ctx := Context()
	sl, err := depl.GetSuggestions(ctx, sr)
	utils.Bail("failed to get suggestions", err)
	for _, sg := range sl.Suggestions {
		fmt.Println(Suggestion2Line(sg))
	}
}
func listSuggestionsOld() {
	ctx := Context()
	started := time.Now()
	depls, err := depl.GetDeploymentsFromCache(ctx, &common.Void{})
	utils.Bail("Failed to get deployments from cache", err)
	fmt.Printf("Got deployments from deploymonkey in %0.1f seconds\n", time.Since(started).Seconds())
	started = time.Now()
	cfg, err := config.GetConfig(depl)
	utils.Bail("Could not get config", err)
	fmt.Printf("Got config from package in %0.1f seconds\n", time.Since(started).Seconds())
	started = time.Now()
	s, err := suggest.Analyse(cfg, depls)
	utils.Bail("Suggestion failed", err)
	fmt.Printf("Got suggestions from package in %0.1f seconds\n", time.Since(started).Seconds())
	fmt.Println(s.String())
	if len(s.Starts) != 0 || len(s.Stops) != 0 {
		os.Exit(1)
	}
}
func applySuggestions() {
	max_tries := 5
	for i := 0; i < max_tries; i++ {
		ctx := Context()
		sl, err := depl.GetSuggestions(ctx, &pb.SuggestRequest{})
		utils.Bail("failed to get suggestions", err)
		if len(sl.Suggestions) == 0 {
			fmt.Printf("No suggestions to apply\n")
			return
		}
		err = try_suggestions(sl)
		if err == nil {
			break
		}
		if err != nil {
			fmt.Printf("Attempt %d of %d - Failed to deploy: %s\n", (i + 1), max_tries, err)
			time.Sleep(time.Duration(5) * time.Second)
		}
	}
	os.Exit(1)
}

func try_suggestions(s *pb.SuggestionList) error {
	var err error
	fmt.Printf("Executing %d requests...\n", len(s.Suggestions))
	for _, start := range s.Suggestions {
		if !start.Start {
			continue
		}
		fmt.Println(Suggestion2Line(start))
		ctx := authremote.ContextWithTimeout(*deploy_timeout)
		fmt.Printf("Deploying %s...\n", start.String())
		d := ToDeployRequest(start)
		_, err = depl.DeployAppOnTarget(ctx, d)
		if err != nil {
			return err
		}

	}
	fmt.Printf("Executing stop requests...\n")
	for _, stop := range s.Suggestions {
		if stop.Start {
			continue
		}
		fmt.Println(Suggestion2Line(stop))
		d := ToUndeployRequest(stop)
		ctx := authremote.ContextWithTimeout(*deploy_timeout)
		fmt.Printf("Undeploying %s...\n", stop.String())
		_, err = depl.UndeployAppOnTarget(ctx, d)
		if err != nil {
			return err
		}
	}
	return nil
}

func listDeployments() {
	ctx := Context()
	depls, err := depl.GetDeploymentsFromCache(ctx, &common.Void{})
	utils.Bail("Failed to get deployments from cache", err)
	fmt.Printf("Current Deployments:\n")
	t := utils.Table{}
	t.AddHeaders("ID", "Instances", "RepoID", "ArtefactID", "Binary", "DeployID")
	for _, d := range depls.Deployments {
		fmt.Printf("%10s: ", d.Host)
		for _, group := range d.Apps {
			for _, app := range group.Applications {
				t.AddUint64(app.ID)
				t.AddUint32(app.Instances)
				t.AddUint64(app.RepositoryID)
				t.AddUint64(app.ArtefactID)
				t.AddString(app.Binary)
				t.AddString(app.DeploymentID)
				t.NewRow()
				fmt.Printf("  %4d x%d ns=%s group=%s repo=%d binary=%s (deploymentid=%s)\n",
					app.ID,
					app.Instances,
					group.Namespace,
					group.GroupID,
					app.RepositoryID,
					app.Binary,
					app.DeploymentID,
				)
			}
		}
	}
	fmt.Println(t.ToPrettyString())
}

func listDeployers() {
	t := &utils.Table{}
	ctx := Context()
	depls, err := depl.GetKnownAutodeployers(ctx, &common.Void{})
	utils.Bail("Failed to get deployers from cache", err)
	t.AddHeaders("IP", "Groups", "GroupCount")
	for _, ad := range depls.Autodeployers {
		t.AddString(ad.IP)
		t.AddString(strings.Join(ad.Groups, " | "))
		t.AddInt(len(ad.Groups))
		t.NewRow()
	}
	fmt.Println(t.ToPrettyString())

}

func delVersion() error {
	ctx := Context()
	vers := *del_version
	_, err := depl.DeleteVersion(ctx, &pb.DelVersionRequest{Version: vers})
	utils.Bail("Failed to get deployers from cache", err)
	return nil
}

func ToDeployRequest(suggestion *pb.Suggestion) *pb.DeployAppRequest {
	return suggestion.DeployRequest

	//	res := &pb.DeployAppRequest{Host: suggestion.Host, AppID: suggestion.App.ID}
	//	return res
}
func ToUndeployRequest(suggestion *pb.Suggestion) *pb.UndeployAppRequest {
	return suggestion.UndeployRequest
	//	res := &pb.UndeployAppRequest{Host:suggestion.Host,DeploymentID: suggestion.}
	//	return res
}
func Suggestion2Line(sg *pb.Suggestion) string {
	s := "stop"
	if sg.Start {
		s = "start"
	}
	return fmt.Sprintf("%s %s on %s", s, sg.App.Binary, sg.Host)
}

func Context() context.Context {
	return authremote.ContextWithTimeout(time.Duration(90) * time.Second)
}
