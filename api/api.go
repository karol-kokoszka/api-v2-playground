package api

import (
	"encoding/json"
	"github.com/scylladb/scylla-cloud/api/gen/siren"
	"net/http"
)

type API struct {
}

func (a *API) GetCloudAccount(w http.ResponseWriter, r *http.Request, accountId int, params siren.GetCloudAccountParams) {
	w.WriteHeader(http.StatusOK)
	res := []siren.CloudAccount{
		{
			CheckStatus:     nil,
			CloudProviderId: nil,
			Id:              nil,
			Owner:           nil,
			Properties:      nil,
			State:           nil,
		},
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *API) CreateCluster(w http.ResponseWriter, r *http.Request, accountId int, params siren.CreateClusterParams) {
	//TODO implement me
	panic("implement me")
}

func (a *API) GetClusterRequestById(w http.ResponseWriter, r *http.Request, accountId int, requestId int, params siren.GetClusterRequestByIdParams) {
	//TODO implement me
	panic("implement me")
}

func (a *API) GetClusterDetails(w http.ResponseWriter, r *http.Request, accountId int, clusterId int, params siren.GetClusterDetailsParams) {
	//TODO implement me
	panic("implement me")
}

func (a *API) GetClusterDCs(w http.ResponseWriter, r *http.Request, accountId int, clusterId int, params siren.GetClusterDCsParams) {
	//TODO implement me
	panic("implement me")
}

func (a *API) DeleteCluster(w http.ResponseWriter, r *http.Request, accountId int, clusterId int, params siren.DeleteClusterParams) {
	//TODO implement me
	panic("implement me")
}

func (a *API) GetFirewallAllowedRules(w http.ResponseWriter, r *http.Request, accountId int, clusterId int, params siren.GetFirewallAllowedRulesParams) {
	//TODO implement me
	panic("implement me")
}

func (a *API) CreateFirewallAllowedRule(w http.ResponseWriter, r *http.Request, accountId int, clusterId int, params siren.CreateFirewallAllowedRuleParams) {
	//TODO implement me
	panic("implement me")
}

func (a *API) DeleteFirewallAllowedRule(w http.ResponseWriter, r *http.Request, accountId int, clusterId int, ruleId int, params siren.DeleteFirewallAllowedRuleParams) {
	//TODO implement me
	panic("implement me")
}

func (a *API) GetVPCPeeringList(w http.ResponseWriter, r *http.Request, accountId int, clusterId int, params siren.GetVPCPeeringListParams) {
	//TODO implement me
	panic("implement me")
}

func (a *API) CreateVPCPeering(w http.ResponseWriter, r *http.Request, accountId int, clusterId int, params siren.CreateVPCPeeringParams) {
	//TODO implement me
	panic("implement me")
}

func (a *API) DeleteVPCPeering(w http.ResponseWriter, r *http.Request, accountId int, clusterId int, peerId int, params siren.DeleteVPCPeeringParams) {
	//TODO implement me
	panic("implement me")
}

func (a *API) GetVPCPeeringInfo(w http.ResponseWriter, r *http.Request, accountId int, clusterId int, peerId int, params siren.GetVPCPeeringInfoParams) {
	//TODO implement me
	panic("implement me")
}

func (a *API) GetClusterNodes(w http.ResponseWriter, r *http.Request, accountId int, clusterId int, params siren.GetClusterNodesParams) {
	//TODO implement me
	panic("implement me")
}

func (a *API) GetClusterRequests(w http.ResponseWriter, r *http.Request, accountId int, clusterId int, params siren.GetClusterRequestsParams) {
	//TODO implement me
	panic("implement me")
}

func (a *API) GetClusters(w http.ResponseWriter, r *http.Request, accountId int, params siren.GetClustersParams) {
	//TODO implement me
	panic("implement me")
}

func (a *API) GetCloudProviderRegion(w http.ResponseWriter, r *http.Request, cloudProviderId int, regionId int, params siren.GetCloudProviderRegionParams) {
	//TODO implement me
	panic("implement me")
}

func (a *API) GetCloudProviderRegions(w http.ResponseWriter, r *http.Request, cloudProviderId int, params siren.GetCloudProviderRegionsParams) {
	//TODO implement me
	panic("implement me")
}

func (a *API) GetCloudProviders(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (a *API) GetScyllaVersions(w http.ResponseWriter, r *http.Request, params siren.GetScyllaVersionsParams) {
	//TODO implement me
	panic("implement me")
}
