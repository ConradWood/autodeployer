package db

/*
 This file was created by mkdb-client.
 The intention is not to modify this file, but you may extend the struct DBAppStats
 in a seperate file (so that you can regenerate this one from time to time)
*/

/*
 PRIMARY KEY: ID
*/

/*
 postgres:
 create sequence appstats_seq;

Main Table:

 CREATE TABLE appstats (id integer primary key default nextval('appstats_seq'),app_id bigint not null  references applicationdefinition (id) on delete cascade  ,ipctrackstatus integer not null  ,unexpectedrestarts integer not null  );

Alter statements:
ALTER TABLE appstats ADD COLUMN IF NOT EXISTS app_id bigint not null references applicationdefinition (id) on delete cascade  default 0;
ALTER TABLE appstats ADD COLUMN IF NOT EXISTS ipctrackstatus integer not null default 0;
ALTER TABLE appstats ADD COLUMN IF NOT EXISTS unexpectedrestarts integer not null default 0;


Archive Table: (structs can be moved from main to archive using Archive() function)

 CREATE TABLE appstats_archive (id integer unique not null,app_id bigint not null,ipctrackstatus integer not null,unexpectedrestarts integer not null);
*/

import (
	"context"
	gosql "database/sql"
	"fmt"
	savepb "golang.conradwood.net/apis/deploymonkey"
	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/go-easyops/sql"
	"os"
	"sync"
)

var (
	default_def_DBAppStats *DBAppStats
)

type DBAppStats struct {
	DB                   *sql.DB
	SQLTablename         string
	SQLArchivetablename  string
	customColumnHandlers []CustomColumnHandler
	lock                 sync.Mutex
}

func init() {
	RegisterDBHandlerFactory(func() Handler {
		return DefaultDBAppStats()
	})
}

func DefaultDBAppStats() *DBAppStats {
	if default_def_DBAppStats != nil {
		return default_def_DBAppStats
	}
	psql, err := sql.Open()
	if err != nil {
		fmt.Printf("Failed to open database: %s\n", err)
		os.Exit(10)
	}
	res := NewDBAppStats(psql)
	ctx := context.Background()
	err = res.CreateTable(ctx)
	if err != nil {
		fmt.Printf("Failed to create table: %s\n", err)
		os.Exit(10)
	}
	default_def_DBAppStats = res
	return res
}
func NewDBAppStats(db *sql.DB) *DBAppStats {
	foo := DBAppStats{DB: db}
	foo.SQLTablename = "appstats"
	foo.SQLArchivetablename = "appstats_archive"
	return &foo
}

func (a *DBAppStats) GetCustomColumnHandlers() []CustomColumnHandler {
	return a.customColumnHandlers
}
func (a *DBAppStats) AddCustomColumnHandler(w CustomColumnHandler) {
	a.lock.Lock()
	a.customColumnHandlers = append(a.customColumnHandlers, w)
	a.lock.Unlock()
}

func (a *DBAppStats) NewQuery() *Query {
	return newQuery(a)
}

// archive. It is NOT transactionally save.
func (a *DBAppStats) Archive(ctx context.Context, id uint64) error {

	// load it
	p, err := a.ByID(ctx, id)
	if err != nil {
		return err
	}

	// now save it to archive:
	_, e := a.DB.ExecContext(ctx, "archive_DBAppStats", "insert into "+a.SQLArchivetablename+" (id,app_id, ipctrackstatus, unexpectedrestarts) values ($1,$2, $3, $4) ", p.ID, p.AppDef.ID, p.IPCTrackStatus, p.UnexpectedRestarts)
	if e != nil {
		return e
	}

	// now delete it.
	a.DeleteByID(ctx, id)
	return nil
}

// return a map with columnname -> value_from_proto
func (a *DBAppStats) buildSaveMap(ctx context.Context, p *savepb.AppStats) (map[string]interface{}, error) {
	extra, err := extraFieldsToStore(ctx, a, p)
	if err != nil {
		return nil, err
	}
	res := make(map[string]interface{})
	res["id"] = a.get_col_from_proto(p, "id")
	res["app_id"] = a.get_col_from_proto(p, "app_id")
	res["ipctrackstatus"] = a.get_col_from_proto(p, "ipctrackstatus")
	res["unexpectedrestarts"] = a.get_col_from_proto(p, "unexpectedrestarts")
	if extra != nil {
		for k, v := range extra {
			res[k] = v
		}
	}
	return res, nil
}

func (a *DBAppStats) Save(ctx context.Context, p *savepb.AppStats) (uint64, error) {
	qn := "save_DBAppStats"
	smap, err := a.buildSaveMap(ctx, p)
	if err != nil {
		return 0, err
	}
	delete(smap, "id") // save without id
	return a.saveMap(ctx, qn, smap, p)
}

// Save using the ID specified
func (a *DBAppStats) SaveWithID(ctx context.Context, p *savepb.AppStats) error {
	qn := "insert_DBAppStats"
	smap, err := a.buildSaveMap(ctx, p)
	if err != nil {
		return err
	}
	_, err = a.saveMap(ctx, qn, smap, p)
	return err
}

// use a hashmap of columnname->values to store to database (see buildSaveMap())
func (a *DBAppStats) saveMap(ctx context.Context, queryname string, smap map[string]interface{}, p *savepb.AppStats) (uint64, error) {
	// Save (and use database default ID generation)

	var rows *gosql.Rows
	var e error

	q_cols := ""
	q_valnames := ""
	q_vals := make([]interface{}, 0)
	deli := ""
	i := 0
	// build the 2 parts of the query (column names and value names) as well as the values themselves
	for colname, val := range smap {
		q_cols = q_cols + deli + colname
		i++
		q_valnames = q_valnames + deli + fmt.Sprintf("$%d", i)
		q_vals = append(q_vals, val)
		deli = ","
	}
	rows, e = a.DB.QueryContext(ctx, queryname, "insert into "+a.SQLTablename+" ("+q_cols+") values ("+q_valnames+") returning id", q_vals...)
	if e != nil {
		return 0, a.Error(ctx, queryname, e)
	}
	defer rows.Close()
	if !rows.Next() {
		return 0, a.Error(ctx, queryname, errors.Errorf("No rows after insert"))
	}
	var id uint64
	e = rows.Scan(&id)
	if e != nil {
		return 0, a.Error(ctx, queryname, errors.Errorf("failed to scan id after insert: %s", e))
	}
	p.ID = id
	return id, nil
}

// if ID==0 save, otherwise update
func (a *DBAppStats) SaveOrUpdate(ctx context.Context, p *savepb.AppStats) error {
	if p.ID == 0 {
		_, err := a.Save(ctx, p)
		return err
	}
	return a.Update(ctx, p)
}
func (a *DBAppStats) Update(ctx context.Context, p *savepb.AppStats) error {
	qn := "DBAppStats_Update"
	_, e := a.DB.ExecContext(ctx, qn, "update "+a.SQLTablename+" set app_id=$1, ipctrackstatus=$2, unexpectedrestarts=$3 where id = $4", a.get_AppDef_ID(p), a.get_IPCTrackStatus(p), a.get_UnexpectedRestarts(p), p.ID)

	return a.Error(ctx, qn, e)
}

// delete by id field
func (a *DBAppStats) DeleteByID(ctx context.Context, p uint64) error {
	qn := "deleteDBAppStats_ByID"
	_, e := a.DB.ExecContext(ctx, qn, "delete from "+a.SQLTablename+" where id = $1", p)
	return a.Error(ctx, qn, e)
}

// get it by primary id
func (a *DBAppStats) ByID(ctx context.Context, p uint64) (*savepb.AppStats, error) {
	qn := "DBAppStats_ByID"
	l, e := a.fromQuery(ctx, qn, "id = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, errors.Errorf("ByID: error scanning (%s)", e))
	}
	if len(l) == 0 {
		return nil, a.Error(ctx, qn, errors.Errorf("No AppStats with id %v", p))
	}
	if len(l) != 1 {
		return nil, a.Error(ctx, qn, errors.Errorf("Multiple (%d) AppStats with id %v", len(l), p))
	}
	return l[0], nil
}

// get it by primary id (nil if no such ID row, but no error either)
func (a *DBAppStats) TryByID(ctx context.Context, p uint64) (*savepb.AppStats, error) {
	qn := "DBAppStats_TryByID"
	l, e := a.fromQuery(ctx, qn, "id = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, errors.Errorf("TryByID: error scanning (%s)", e))
	}
	if len(l) == 0 {
		return nil, nil
	}
	if len(l) != 1 {
		return nil, a.Error(ctx, qn, errors.Errorf("Multiple (%d) AppStats with id %v", len(l), p))
	}
	return l[0], nil
}

// get it by multiple primary ids
func (a *DBAppStats) ByIDs(ctx context.Context, p []uint64) ([]*savepb.AppStats, error) {
	qn := "DBAppStats_ByIDs"
	l, e := a.fromQuery(ctx, qn, "id in $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, errors.Errorf("TryByID: error scanning (%s)", e))
	}
	return l, nil
}

// get all rows
func (a *DBAppStats) All(ctx context.Context) ([]*savepb.AppStats, error) {
	qn := "DBAppStats_all"
	l, e := a.fromQuery(ctx, qn, "true")
	if e != nil {
		return nil, errors.Errorf("All: error scanning (%s)", e)
	}
	return l, nil
}

/**********************************************************************
* GetBy[FIELD] functions
**********************************************************************/

// get all "DBAppStats" rows with matching AppDef
func (a *DBAppStats) ByAppDef(ctx context.Context, p uint64) ([]*savepb.AppStats, error) {
	qn := "DBAppStats_ByAppDef"
	l, e := a.fromQuery(ctx, qn, "app_id = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, errors.Errorf("ByAppDef: error scanning (%s)", e))
	}
	return l, nil
}

// get all "DBAppStats" rows with multiple matching AppDef
func (a *DBAppStats) ByMultiAppDef(ctx context.Context, p []uint64) ([]*savepb.AppStats, error) {
	qn := "DBAppStats_ByAppDef"
	l, e := a.fromQuery(ctx, qn, "app_id in $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, errors.Errorf("ByAppDef: error scanning (%s)", e))
	}
	return l, nil
}

// the 'like' lookup
func (a *DBAppStats) ByLikeAppDef(ctx context.Context, p uint64) ([]*savepb.AppStats, error) {
	qn := "DBAppStats_ByLikeAppDef"
	l, e := a.fromQuery(ctx, qn, "app_id ilike $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, errors.Errorf("ByAppDef: error scanning (%s)", e))
	}
	return l, nil
}

// get all "DBAppStats" rows with matching IPCTrackStatus
func (a *DBAppStats) ByIPCTrackStatus(ctx context.Context, p uint32) ([]*savepb.AppStats, error) {
	qn := "DBAppStats_ByIPCTrackStatus"
	l, e := a.fromQuery(ctx, qn, "ipctrackstatus = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, errors.Errorf("ByIPCTrackStatus: error scanning (%s)", e))
	}
	return l, nil
}

// get all "DBAppStats" rows with multiple matching IPCTrackStatus
func (a *DBAppStats) ByMultiIPCTrackStatus(ctx context.Context, p []uint32) ([]*savepb.AppStats, error) {
	qn := "DBAppStats_ByIPCTrackStatus"
	l, e := a.fromQuery(ctx, qn, "ipctrackstatus in $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, errors.Errorf("ByIPCTrackStatus: error scanning (%s)", e))
	}
	return l, nil
}

// the 'like' lookup
func (a *DBAppStats) ByLikeIPCTrackStatus(ctx context.Context, p uint32) ([]*savepb.AppStats, error) {
	qn := "DBAppStats_ByLikeIPCTrackStatus"
	l, e := a.fromQuery(ctx, qn, "ipctrackstatus ilike $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, errors.Errorf("ByIPCTrackStatus: error scanning (%s)", e))
	}
	return l, nil
}

// get all "DBAppStats" rows with matching UnexpectedRestarts
func (a *DBAppStats) ByUnexpectedRestarts(ctx context.Context, p uint32) ([]*savepb.AppStats, error) {
	qn := "DBAppStats_ByUnexpectedRestarts"
	l, e := a.fromQuery(ctx, qn, "unexpectedrestarts = $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, errors.Errorf("ByUnexpectedRestarts: error scanning (%s)", e))
	}
	return l, nil
}

// get all "DBAppStats" rows with multiple matching UnexpectedRestarts
func (a *DBAppStats) ByMultiUnexpectedRestarts(ctx context.Context, p []uint32) ([]*savepb.AppStats, error) {
	qn := "DBAppStats_ByUnexpectedRestarts"
	l, e := a.fromQuery(ctx, qn, "unexpectedrestarts in $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, errors.Errorf("ByUnexpectedRestarts: error scanning (%s)", e))
	}
	return l, nil
}

// the 'like' lookup
func (a *DBAppStats) ByLikeUnexpectedRestarts(ctx context.Context, p uint32) ([]*savepb.AppStats, error) {
	qn := "DBAppStats_ByLikeUnexpectedRestarts"
	l, e := a.fromQuery(ctx, qn, "unexpectedrestarts ilike $1", p)
	if e != nil {
		return nil, a.Error(ctx, qn, errors.Errorf("ByUnexpectedRestarts: error scanning (%s)", e))
	}
	return l, nil
}

/**********************************************************************
* The field getters
**********************************************************************/

// getter for field "ID" (ID) [uint64]
func (a *DBAppStats) get_ID(p *savepb.AppStats) uint64 {
	return uint64(p.ID)
}

// getter for reference "AppDef"
func (a *DBAppStats) get_AppDef_ID(p *savepb.AppStats) uint64 {
	if p.AppDef == nil {
		panic("field AppDef must not be nil")
	}
	return p.AppDef.ID
}

// getter for field "IPCTrackStatus" (IPCTrackStatus) [uint32]
func (a *DBAppStats) get_IPCTrackStatus(p *savepb.AppStats) uint32 {
	return uint32(p.IPCTrackStatus)
}

// getter for field "UnexpectedRestarts" (UnexpectedRestarts) [uint32]
func (a *DBAppStats) get_UnexpectedRestarts(p *savepb.AppStats) uint32 {
	return uint32(p.UnexpectedRestarts)
}

/**********************************************************************
* Helper to convert from an SQL Query
**********************************************************************/

// from a query snippet (the part after WHERE)
func (a *DBAppStats) ByDBQuery(ctx context.Context, query *Query) ([]*savepb.AppStats, error) {
	extra_fields, err := extraFieldsToQuery(ctx, a)
	if err != nil {
		return nil, err
	}
	i := 0
	for col_name, value := range extra_fields {
		i++
		/*
		   efname:=fmt.Sprintf("EXTRA_FIELD_%d",i)
		   query.Add(col_name+" = "+efname,QP{efname:value})
		*/
		query.AddEqual(col_name, value)
	}

	gw, paras := query.ToPostgres()
	queryname := "custom_dbquery"
	rows, err := a.DB.QueryContext(ctx, queryname, "select "+a.SelectCols()+" from "+a.Tablename()+" where "+gw, paras...)
	if err != nil {
		return nil, err
	}
	res, err := a.FromRows(ctx, rows)
	rows.Close()
	if err != nil {
		return nil, err
	}
	return res, nil

}

func (a *DBAppStats) FromQuery(ctx context.Context, query_where string, args ...interface{}) ([]*savepb.AppStats, error) {
	return a.fromQuery(ctx, "custom_query_"+a.Tablename(), query_where, args...)
}

// from a query snippet (the part after WHERE)
func (a *DBAppStats) fromQuery(ctx context.Context, queryname string, query_where string, args ...interface{}) ([]*savepb.AppStats, error) {
	extra_fields, err := extraFieldsToQuery(ctx, a)
	if err != nil {
		return nil, err
	}
	eq := ""
	if extra_fields != nil && len(extra_fields) > 0 {
		eq = " AND ("
		// build the extraquery "eq"
		i := len(args)
		deli := ""
		for col_name, value := range extra_fields {
			i++
			eq = eq + deli + col_name + fmt.Sprintf(" = $%d", i)
			deli = " AND "
			args = append(args, value)
		}
		eq = eq + ")"
	}
	rows, err := a.DB.QueryContext(ctx, queryname, "select "+a.SelectCols()+" from "+a.Tablename()+" where ( "+query_where+") "+eq, args...)
	if err != nil {
		return nil, err
	}
	res, err := a.FromRows(ctx, rows)
	rows.Close()
	if err != nil {
		return nil, err
	}
	return res, nil
}

/**********************************************************************
* Helper to convert from an SQL Row to struct
**********************************************************************/
func (a *DBAppStats) get_col_from_proto(p *savepb.AppStats, colname string) interface{} {
	if colname == "id" {
		return a.get_ID(p)
	} else if colname == "app_id" {
		return a.get_AppDef_ID(p)
	} else if colname == "ipctrackstatus" {
		return a.get_IPCTrackStatus(p)
	} else if colname == "unexpectedrestarts" {
		return a.get_UnexpectedRestarts(p)
	}
	panic(fmt.Sprintf("in table \"%s\", column \"%s\" cannot be resolved to proto field name", a.Tablename(), colname))
}

func (a *DBAppStats) Tablename() string {
	return a.SQLTablename
}

func (a *DBAppStats) SelectCols() string {
	return "id,app_id, ipctrackstatus, unexpectedrestarts"
}
func (a *DBAppStats) SelectColsQualified() string {
	return "" + a.SQLTablename + ".id," + a.SQLTablename + ".app_id, " + a.SQLTablename + ".ipctrackstatus, " + a.SQLTablename + ".unexpectedrestarts"
}

func (a *DBAppStats) FromRows(ctx context.Context, rows *gosql.Rows) ([]*savepb.AppStats, error) {
	var res []*savepb.AppStats
	for rows.Next() {
		// SCANNER:
		foo := &savepb.AppStats{}
		// create the non-nullable pointers
		foo.AppDef = &savepb.ApplicationDefinition{} // non-nullable
		// create variables for scan results
		scanTarget_0 := &foo.ID
		scanTarget_1 := &foo.AppDef.ID
		scanTarget_2 := &foo.IPCTrackStatus
		scanTarget_3 := &foo.UnexpectedRestarts
		err := rows.Scan(scanTarget_0, scanTarget_1, scanTarget_2, scanTarget_3)
		// END SCANNER

		if err != nil {
			return nil, a.Error(ctx, "fromrow-scan", err)
		}
		res = append(res, foo)
	}
	return res, nil
}

/**********************************************************************
* Helper to create table and columns
**********************************************************************/
func (a *DBAppStats) CreateTable(ctx context.Context) error {
	csql := []string{
		`create sequence if not exists ` + a.SQLTablename + `_seq;`,
		`CREATE TABLE if not exists ` + a.SQLTablename + ` (id integer primary key default nextval('` + a.SQLTablename + `_seq'),app_id bigint not null ,ipctrackstatus integer not null ,unexpectedrestarts integer not null );`,
		`CREATE TABLE if not exists ` + a.SQLTablename + `_archive (id integer primary key default nextval('` + a.SQLTablename + `_seq'),app_id bigint not null ,ipctrackstatus integer not null ,unexpectedrestarts integer not null );`,
		`ALTER TABLE ` + a.SQLTablename + ` ADD COLUMN IF NOT EXISTS app_id bigint not null default 0;`,
		`ALTER TABLE ` + a.SQLTablename + ` ADD COLUMN IF NOT EXISTS ipctrackstatus integer not null default 0;`,
		`ALTER TABLE ` + a.SQLTablename + ` ADD COLUMN IF NOT EXISTS unexpectedrestarts integer not null default 0;`,

		`ALTER TABLE ` + a.SQLTablename + `_archive  ADD COLUMN IF NOT EXISTS app_id bigint not null  default 0;`,
		`ALTER TABLE ` + a.SQLTablename + `_archive  ADD COLUMN IF NOT EXISTS ipctrackstatus integer not null  default 0;`,
		`ALTER TABLE ` + a.SQLTablename + `_archive  ADD COLUMN IF NOT EXISTS unexpectedrestarts integer not null  default 0;`,
	}

	for i, c := range csql {
		_, e := a.DB.ExecContext(ctx, fmt.Sprintf("create_"+a.SQLTablename+"_%d", i), c)
		if e != nil {
			return e
		}
	}

	// these are optional, expected to fail
	csql = []string{
		// Indices:

		// Foreign keys:
		`ALTER TABLE ` + a.SQLTablename + ` add constraint mkdb_fk_appstats_app_id_applicationdefinitionid FOREIGN KEY (app_id) references applicationdefinition (id) on delete cascade ;`,
	}
	for i, c := range csql {
		a.DB.ExecContextQuiet(ctx, fmt.Sprintf("create_"+a.SQLTablename+"_%d", i), c)
	}
	return nil
}

/**********************************************************************
* Helper to meaningful errors
**********************************************************************/
func (a *DBAppStats) Error(ctx context.Context, q string, e error) error {
	if e == nil {
		return nil
	}
	return errors.Errorf("[table="+a.SQLTablename+", query=%s] Error: %s", q, e)
}

