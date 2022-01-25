package main

import (
	"fmt"
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/openstack"
	"github.com/huaweicloud/golangsdk/openstack/compute/v2/extensions/aggregates"
	"github.com/huaweicloud/golangsdk/openstack/compute/v2/extensions/availabilityzones"
	"github.com/huaweicloud/golangsdk/openstack/compute/v2/extensions/hypervisors"
	"github.com/huaweicloud/golangsdk/openstack/compute/v2/flavors"
	"github.com/huaweicloud/golangsdk/openstack/compute/v2/images"
	"github.com/huaweicloud/golangsdk/openstack/compute/v2/servers"   //for servers
	"github.com/huaweicloud/golangsdk/openstack/identity/v2/tenants"  //for v2租户
	"github.com/huaweicloud/golangsdk/openstack/identity/v3/projects" //for v3租户
	"github.com/huaweicloud/golangsdk/openstack/networking/v1/subnets"
	"github.com/huaweicloud/golangsdk/openstack/networking/v2/extensions/layer3/floatingips"
	"github.com/huaweicloud/golangsdk/openstack/networking/v2/extensions/layer3/routers"
	"github.com/huaweicloud/golangsdk/openstack/networking/v2/extensions/security/groups"
	"github.com/huaweicloud/golangsdk/openstack/networking/v2/networks"
	"github.com/huaweicloud/golangsdk/openstack/networking/v2/ports"
	_ "github.com/huaweicloud/golangsdk/pagination" //for EachPage
	"log"
	//"github.com/huaweicloud/golangsdk"
	//"github.com/huaweicloud/golangsdk/openstack"
	//"github.com/huaweicloud/golangsdk/openstack/utils"
)

func getNetworksVpc(client *golangsdk.ServiceClient) bool {
	allpages, err := networks.List(client, networks.ListOpts{}).AllPages()
	if err != nil {
		panic(err)
	}

	info, err := networks.ExtractNetworks(allpages)
	if err != nil {
		panic(info)
	}
	if len(info) > 0 {
		fmt.Printf("networks:%v\n", info)
		return true
	}

	return false
}

func getSubnets(client *golangsdk.ServiceClient) bool {
	info, err := subnets.List(client, subnets.ListOpts{})
	if err != nil {
		panic(err)
	}

	if len(info) > 0 {
		fmt.Printf("subnets:%v\n", info)
		return true
	}

	return false
}

func getVports(client *golangsdk.ServiceClient) bool {
	allpages, err := ports.List(client, ports.ListOpts{}).AllPages()
	if err != nil {
		panic(err)
	}

	info, err := ports.ExtractPorts(allpages)
	if err != nil {
		panic(info)
	}
	if len(info) > 0 {
		fmt.Printf("vports:%v\n", info)
		return true
	}

	return false
}

func getFloatingips(client *golangsdk.ServiceClient) bool {
	allpages, err := floatingips.List(client, floatingips.ListOpts{}).AllPages()
	if err != nil {
		panic(err)
	}

	info, err := floatingips.ExtractFloatingIPs(allpages)
	if err != nil {
		panic(info)
	}
	if len(info) > 0 {
		fmt.Printf("floatingips:%v\n", info)
		return true
	}

	return false
}

func getRouters(client *golangsdk.ServiceClient) bool {
	allpages, err := routers.List(client, routers.ListOpts{}).AllPages()
	if err != nil {
		panic(err)
	}

	info, err := routers.ExtractRouters(allpages)
	if err != nil {
		panic(info)
	}
	if len(info) > 0 {
		fmt.Printf("routers:%v\n", info)
		return true
	}

	return false
}

func getSecurityGroup(client *golangsdk.ServiceClient) bool {
	allpages, err := groups.List(client, groups.ListOpts{}).AllPages()
	if err != nil {
		panic(err)
	}

	info, err := groups.ExtractGroups(allpages)
	if err != nil {
		panic(info)
	}
	if len(info) > 0 {
		fmt.Printf("SecurityGroup:%v\n", info)
		return true

	}

	return false
}

func getSecurityGroupById(client *golangsdk.ServiceClient, id string) bool {
	r := groups.Get(client, id)
	secGroup, err := r.Extract()
	if err != nil {
		panic(err)
	}

	if secGroup != nil {
		fmt.Printf("SecurityGroup of %s:%v\n", id, secGroup)
		return true
	}
	return false
}

//网络资源
func networkApi(client *golangsdk.ServiceClient) {
	getNetworksVpc(client)
	getSubnets(client)
	getVports(client)
	getFloatingips(client)
	getRouters(client)
	getSecurityGroup(client)

	secGroupId := "123456"
	getSecurityGroupById(client, secGroupId)
}

//获取虚拟主机信息
func getVHostInfo(client *golangsdk.ServiceClient) bool {
	listOpts := servers.ListOpts{
		AllTenants: true, //TODO:还有很多选项
	}

	allPages, err := servers.List(client, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allServers, err := servers.ExtractServers(allPages)
	if err != nil {
		panic(err)
	}

	for _, server := range allServers {
		fmt.Printf("%+v\n", server)
	}

	//也可以使用下面方式
	/*pages := 0
	err = servers.List(client.ServiceClient(), servers.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		pages++

		_, err := servers.ExtractServers(page)
		if err != nil {
			return false, err
		}

		return true, nil
	})

	fmt.Printf("servers.List.EachPage error:%s\n", err.Error())*/

	return true
}

//获取虚拟主机信息
func getHypervisorsInfo(client *golangsdk.ServiceClient) bool {
	allpages, err := hypervisors.List(client).AllPages()
	if err != nil {
		panic(err)
	}

	Hypervisors, err := hypervisors.ExtractHypervisors(allpages)
	if err != nil {
		panic(err)
	}

	fmt.Printf("hypervisors:%v\n", Hypervisors)

	return true
}

//获取主机池信息
func getHostPoolInfo(client *golangsdk.ServiceClient) bool {
	allPages, err := availabilityzones.ListDetail(client).AllPages()
	if err != nil {
		panic(err)
	}

	azs, err := availabilityzones.ExtractAvailabilityZones(allPages)
	if err != nil {
		panic(err)
	}

	if len(azs) > 0 {
		fmt.Printf("AvailabilityZones:%v\n", azs)
		return true
	}

	return false
}

//获取集群信息
func getAggregatesInfo(client *golangsdk.ServiceClient) bool {
	allpages, err := aggregates.List(client).AllPages()
	if err != nil {
		panic(err)
	}

	aggs, err := aggregates.ExtractAggregates(allpages)
	if err != nil {
		panic(err)
	}

	if len(aggs) > 0 {
		fmt.Printf("Aggregates:%v\n", aggs)
		return true
	}

	fmt.Println("Aggregates record empty")
	return false
}

func getFlavors(client *golangsdk.ServiceClient) bool {
	allpages, err := flavors.ListDetail(client, flavors.ListOpts{}).AllPages()
	if err != nil {
		panic(err)
	}

	info, err := flavors.ExtractFlavors(allpages)
	if err != nil {
		panic(err)
	}

	if len(info) > 0 {
		fmt.Printf("flavors:%v\n", info)
		return true
	}

	fmt.Println("flavors record empty")
	return false
}

func getImages(client *golangsdk.ServiceClient) bool {
	allpages, err := images.ListDetail(client, images.ListOpts{}).AllPages()
	if err != nil {
		panic(err)
	}

	info, err := images.ExtractImages(allpages)
	if err != nil {
		panic(err)
	}

	if len(info) > 0 {
		fmt.Printf("images:%v\n", info)
		return true
	}

	fmt.Println("images record empty")
	return false
}

//计算资源
func computeApi(client *golangsdk.ServiceClient) {
	getVHostInfo(client)

	getHypervisorsInfo(client)

	getAggregatesInfo(client)

	getHostPoolInfo(client)

	getFlavors(client)

	getImages(client)
}

//获取租户资源 version 2版本
func identifyApiV2(client *golangsdk.ServiceClient) bool {
	allpages, err := tenants.List(client, &tenants.ListOpts{}).AllPages()
	if err != nil {
		panic(err)
	}

	tenantsInfo, err := tenants.ExtractTenants(allpages)
	if err != nil {
		panic(err)
	}

	if len(tenantsInfo) > 0 {
		fmt.Printf("tenants:%v\n", tenantsInfo)
		return true
	}

	return false
}

//获取租户资源 version 3版本
func identifyApiV3(client *golangsdk.ServiceClient) bool {
	allpages, err := projects.List(client, &projects.ListOpts{}).AllPages()
	if err != nil {
		panic(err)
	}

	projectsInfo, err := projects.ExtractProjects(allpages)
	if err != nil {
		panic(err)
	}

	if len(projectsInfo) > 0 {
		fmt.Printf("tenants:%v\n", projectsInfo)
		return true
	}

	return false
}

//补充
func main() {
	//登录IAM服务器

	//用户名、密码、节点url是必填项
	opts := golangsdk.AuthOptions{
		IdentityEndpoint: "https://openstack.example.com:5000/v2.0", //TODO:
		Username:         "{username}",                              //TODO:
		Password:         "{password}",                              //TODO:
		TenantName:       "admin",
	}

	//另一种方式是读取环境变量，暂时没采用
	//opts2, err := openstack.AuthOptionsFromEnv()
	provider, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		log.Fatal("golangsdk.AuthenticatedClient create client failed!")
	}

	///////////////////////////////////计算资源测试////////////////////////////////////////////
	computeClient, err := openstack.NewComputeV2(provider, golangsdk.EndpointOpts{
		Region: "xxx", //TODO:
	})
	if err != nil {
		log.Fatal("openstack.NewComputeV2 create compute client failed")
	}
	computeApi(computeClient)

	///////////////////////////////////网络资源测试////////////////////////////////////////////
	networkClient, err := openstack.NewNetworkV2(provider, golangsdk.EndpointOpts{
		Region: "xxx", //TODO:
	})
	if err != nil {
		log.Fatal("openstack.NewComputeV2 create compute client failed")
	}
	networkApi(networkClient)

	///////////////////////////////////Identify资源测试////////////////////////////////////////////
	//version 2测试
	identifyClientV2, err := openstack.NewIdentityV2(provider, golangsdk.EndpointOpts{
		Region: "xxx", //TODO:
	})
	identifyApiV2(identifyClientV2)

	//version 3测试
	identifyClientV3, err := openstack.NewIdentityV3(provider, golangsdk.EndpointOpts{
		Region: "xxx", //TODO:
	})
	identifyApiV3(identifyClientV3)
}
