// client create: EchoStreamServiceClient
/* geninfo:
   filename  : golang.conradwood.net/apis/openweather/openweather.proto
   gopackage : golang.conradwood.net/apis/openweather
   importname: ai_1
   varname   : client_EchoStreamServiceClient_1
   clientname: EchoStreamServiceClient
   servername: EchoStreamServiceServer
   gscvname  : openweather.EchoStreamService
   lockname  : lock_EchoStreamServiceClient_1
   activename: active_EchoStreamServiceClient_1
*/

package openweather

import (
   "sync"
   "golang.conradwood.net/go-easyops/client"
)
var (
  lock_EchoStreamServiceClient_1 sync.Mutex
  client_EchoStreamServiceClient_1 EchoStreamServiceClient
)

func GetEchoStreamClient() EchoStreamServiceClient { 
    if client_EchoStreamServiceClient_1 != nil {
        return client_EchoStreamServiceClient_1
    }

    lock_EchoStreamServiceClient_1.Lock() 
    if client_EchoStreamServiceClient_1 != nil {
       lock_EchoStreamServiceClient_1.Unlock()
       return client_EchoStreamServiceClient_1
    }

    client_EchoStreamServiceClient_1 = NewEchoStreamServiceClient(client.Connect("openweather.EchoStreamService"))
    lock_EchoStreamServiceClient_1.Unlock()
    return client_EchoStreamServiceClient_1
}

func GetEchoStreamServiceClient() EchoStreamServiceClient { 
    if client_EchoStreamServiceClient_1 != nil {
        return client_EchoStreamServiceClient_1
    }

    lock_EchoStreamServiceClient_1.Lock() 
    if client_EchoStreamServiceClient_1 != nil {
       lock_EchoStreamServiceClient_1.Unlock()
       return client_EchoStreamServiceClient_1
    }

    client_EchoStreamServiceClient_1 = NewEchoStreamServiceClient(client.Connect("openweather.EchoStreamService"))
    lock_EchoStreamServiceClient_1.Unlock()
    return client_EchoStreamServiceClient_1
}

