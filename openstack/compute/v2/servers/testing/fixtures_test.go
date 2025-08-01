package testing

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/servers"
	th "github.com/gophercloud/gophercloud/v2/testhelper"
	"github.com/gophercloud/gophercloud/v2/testhelper/client"
)

// ServerListBody contains the canned body of a servers.List response.
const ServerListBody = `
{
	"servers": [
		{
			"status": "ACTIVE",
			"updated": "2014-09-25T13:10:10Z",
			"hostId": "29d3c8c896a45aa4c34e52247875d7fefc3d94bbcc9f622b5d204362",
			"OS-EXT-SRV-ATTR:host": "devstack",
			"addresses": {
				"private": [
					{
						"OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:7c:1b:2b",
						"version": 4,
						"addr": "10.0.0.32",
						"OS-EXT-IPS:type": "fixed"
					}
				]
			},
			"links": [
				{
					"href": "http://104.130.131.164:8774/v2/fcad67a6189847c4aecfa3c81a05783b/servers/ef079b0c-e610-4dfb-b1aa-b49f07ac48e5",
					"rel": "self"
				},
				{
					"href": "http://104.130.131.164:8774/fcad67a6189847c4aecfa3c81a05783b/servers/ef079b0c-e610-4dfb-b1aa-b49f07ac48e5",
					"rel": "bookmark"
				}
			],
			"key_name": null,
			"image": {
				"id": "f90f6034-2570-4974-8351-6b49732ef2eb",
				"links": [
					{
						"href": "http://104.130.131.164:8774/fcad67a6189847c4aecfa3c81a05783b/images/f90f6034-2570-4974-8351-6b49732ef2eb",
						"rel": "bookmark"
					}
				]
			},
			"OS-EXT-STS:task_state": null,
			"OS-EXT-STS:vm_state": "active",
			"OS-EXT-SRV-ATTR:instance_name": "instance-0000001e",
			"OS-SRV-USG:launched_at": "2014-09-25T13:10:10.000000",
			"OS-EXT-SRV-ATTR:hypervisor_hostname": "devstack",
			"flavor": {
				"id": "1",
				"links": [
					{
						"href": "http://104.130.131.164:8774/fcad67a6189847c4aecfa3c81a05783b/flavors/1",
						"rel": "bookmark"
					}
				]
			},
			"id": "ef079b0c-e610-4dfb-b1aa-b49f07ac48e5",
			"security_groups": [
				{
					"name": "default"
				}
			],
			"OS-SRV-USG:terminated_at": null,
			"OS-EXT-AZ:availability_zone": "nova",
			"user_id": "9349aff8be7545ac9d2f1d00999a23cd",
			"name": "herp",
			"created": "2014-09-25T13:10:02Z",
			"tenant_id": "fcad67a6189847c4aecfa3c81a05783b",
			"OS-DCF:diskConfig": "MANUAL",
			"os-extended-volumes:volumes_attached": [
				{
					"id": "2bdbc40f-a277-45d4-94ac-d9881c777d33"
				}
			],
			"accessIPv4": "",
			"accessIPv6": "",
			"progress": 0,
			"OS-EXT-STS:power_state": 1,
			"config_drive": "",
			"metadata": {}
		},
		{
			"status": "ACTIVE",
			"updated": "2014-09-25T13:04:49Z",
			"hostId": "29d3c8c896a45aa4c34e52247875d7fefc3d94bbcc9f622b5d204362",
			"OS-EXT-SRV-ATTR:host": "devstack",
			"addresses": {
				"private": [
					{
						"OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:9e:89:be",
						"version": 4,
						"addr": "10.0.0.31",
						"OS-EXT-IPS:type": "fixed"
					}
				]
			},
			"links": [
				{
					"href": "http://104.130.131.164:8774/v2/fcad67a6189847c4aecfa3c81a05783b/servers/9e5476bd-a4ec-4653-93d6-72c93aa682ba",
					"rel": "self"
				},
				{
					"href": "http://104.130.131.164:8774/fcad67a6189847c4aecfa3c81a05783b/servers/9e5476bd-a4ec-4653-93d6-72c93aa682ba",
					"rel": "bookmark"
				}
			],
			"key_name": null,
			"image": {
				"id": "f90f6034-2570-4974-8351-6b49732ef2eb",
				"links": [
					{
						"href": "http://104.130.131.164:8774/fcad67a6189847c4aecfa3c81a05783b/images/f90f6034-2570-4974-8351-6b49732ef2eb",
						"rel": "bookmark"
					}
				]
			},
			"OS-EXT-STS:task_state": null,
			"OS-EXT-STS:vm_state": "active",
			"OS-EXT-SRV-ATTR:instance_name": "instance-0000001d",
			"OS-SRV-USG:launched_at": "2014-09-25T13:04:49.000000",
			"OS-EXT-SRV-ATTR:hypervisor_hostname": "devstack",
			"flavor": {
				"id": "1",
				"links": [
					{
						"href": "http://104.130.131.164:8774/fcad67a6189847c4aecfa3c81a05783b/flavors/1",
						"rel": "bookmark"
					}
				]
			},
			"id": "9e5476bd-a4ec-4653-93d6-72c93aa682ba",
			"security_groups": [
				{
					"name": "default"
				}
			],
			"OS-SRV-USG:terminated_at": null,
			"OS-EXT-AZ:availability_zone": "nova",
			"user_id": "9349aff8be7545ac9d2f1d00999a23cd",
			"name": "derp",
			"created": "2014-09-25T13:04:41Z",
			"tenant_id": "fcad67a6189847c4aecfa3c81a05783b",
			"OS-DCF:diskConfig": "MANUAL",
			"os-extended-volumes:volumes_attached": [],
			"accessIPv4": "",
			"accessIPv6": "",
			"progress": 0,
			"OS-EXT-STS:power_state": 1,
			"config_drive": "",
			"metadata": {},
			"locked": true
		},
		{
		"status": "ACTIVE",
		"updated": "2014-09-25T13:04:49Z",
		"hostId": "29d3c8c896a45aa4c34e52247875d7fefc3d94bbcc9f622b5d204362",
		"OS-EXT-SRV-ATTR:host": "devstack",
		"addresses": {
			"private": [
				{
					"OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:9e:89:be",
					"version": 4,
					"addr": "10.0.0.31",
					"OS-EXT-IPS:type": "fixed"
				}
			]
		},
		"links": [
			{
				"href": "http://104.130.131.164:8774/v2/fcad67a6189847c4aecfa3c81a05783b/servers/9e5476bd-a4ec-4653-93d6-72c93aa682ba",
				"rel": "self"
			},
			{
				"href": "http://104.130.131.164:8774/fcad67a6189847c4aecfa3c81a05783b/servers/9e5476bd-a4ec-4653-93d6-72c93aa682ba",
				"rel": "bookmark"
			}
		],
		"key_name": null,
		"image": "",
		"OS-EXT-STS:task_state": null,
		"OS-EXT-STS:vm_state": "active",
		"OS-EXT-SRV-ATTR:instance_name": "instance-0000001d",
		"OS-SRV-USG:launched_at": "2014-09-25T13:04:49.000000",
		"OS-EXT-SRV-ATTR:hypervisor_hostname": "devstack",
		"flavor": {
			"id": "1",
			"links": [
				{
					"href": "http://104.130.131.164:8774/fcad67a6189847c4aecfa3c81a05783b/flavors/1",
					"rel": "bookmark"
				}
			]
		},
		"id": "9e5476bd-a4ec-4653-93d6-72c93aa682bb",
		"security_groups": [
			{
				"name": "default"
			}
		],
		"OS-SRV-USG:terminated_at": null,
		"OS-EXT-AZ:availability_zone": "nova",
		"user_id": "9349aff8be7545ac9d2f1d00999a23cd",
		"name": "merp",
		"created": "2014-09-25T13:04:41Z",
		"tenant_id": "fcad67a6189847c4aecfa3c81a05783b",
		"OS-DCF:diskConfig": "MANUAL",
		"os-extended-volumes:volumes_attached": [],
		"accessIPv4": "",
		"accessIPv6": "",
		"progress": 0,
		"OS-EXT-STS:power_state": 1,
		"config_drive": "",
		"metadata": {}
	}
	]
}
`

// SingleServerBody is the canned body of a Get request on an existing server.
const SingleServerBody = `
{
	"server": {
		"status": "ACTIVE",
		"updated": "2014-09-25T13:04:49Z",
		"hostId": "29d3c8c896a45aa4c34e52247875d7fefc3d94bbcc9f622b5d204362",
		"OS-EXT-SRV-ATTR:host": "devstack",
		"addresses": {
			"private": [
				{
					"OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:9e:89:be",
					"version": 4,
					"addr": "10.0.0.31",
					"OS-EXT-IPS:type": "fixed"
				}
			]
		},
		"links": [
			{
				"href": "http://104.130.131.164:8774/v2/fcad67a6189847c4aecfa3c81a05783b/servers/9e5476bd-a4ec-4653-93d6-72c93aa682ba",
				"rel": "self"
			},
			{
				"href": "http://104.130.131.164:8774/fcad67a6189847c4aecfa3c81a05783b/servers/9e5476bd-a4ec-4653-93d6-72c93aa682ba",
				"rel": "bookmark"
			}
		],
		"key_name": null,
		"image": {
			"id": "f90f6034-2570-4974-8351-6b49732ef2eb",
			"links": [
				{
					"href": "http://104.130.131.164:8774/fcad67a6189847c4aecfa3c81a05783b/images/f90f6034-2570-4974-8351-6b49732ef2eb",
					"rel": "bookmark"
				}
			]
		},
		"OS-EXT-STS:task_state": null,
		"OS-EXT-STS:vm_state": "active",
		"OS-EXT-SRV-ATTR:instance_name": "instance-0000001d",
		"OS-SRV-USG:launched_at": "2014-09-25T13:04:49.000000",
		"OS-EXT-SRV-ATTR:hypervisor_hostname": "devstack",
		"flavor": {
			"id": "1",
			"links": [
				{
					"href": "http://104.130.131.164:8774/fcad67a6189847c4aecfa3c81a05783b/flavors/1",
					"rel": "bookmark"
				}
			]
		},
		"id": "9e5476bd-a4ec-4653-93d6-72c93aa682ba",
		"security_groups": [
			{
				"name": "default"
			}
		],
		"OS-SRV-USG:terminated_at": null,
		"OS-EXT-AZ:availability_zone": "nova",
		"user_id": "9349aff8be7545ac9d2f1d00999a23cd",
		"name": "derp",
		"created": "2014-09-25T13:04:41Z",
		"tenant_id": "fcad67a6189847c4aecfa3c81a05783b",
		"OS-DCF:diskConfig": "MANUAL",
		"os-extended-volumes:volumes_attached": [],
		"accessIPv4": "",
		"accessIPv6": "",
		"progress": 0,
		"OS-EXT-STS:power_state": 1,
		"config_drive": "",
		"metadata": {},
		"locked": true
	}
}
`

// FaultyServerBody is the body of a Get request on an existing server
// which has a fault/error.
const FaultyServerBody = `
{
	"server": {
		"status": "ACTIVE",
		"updated": "2014-09-25T13:04:49Z",
		"hostId": "29d3c8c896a45aa4c34e52247875d7fefc3d94bbcc9f622b5d204362",
		"OS-EXT-SRV-ATTR:host": "devstack",
		"addresses": {
			"private": [
				{
					"OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:9e:89:be",
					"version": 4,
					"addr": "10.0.0.31",
					"OS-EXT-IPS:type": "fixed"
				}
			]
		},
		"links": [
			{
				"href": "http://104.130.131.164:8774/v2/fcad67a6189847c4aecfa3c81a05783b/servers/9e5476bd-a4ec-4653-93d6-72c93aa682ba",
				"rel": "self"
			},
			{
				"href": "http://104.130.131.164:8774/fcad67a6189847c4aecfa3c81a05783b/servers/9e5476bd-a4ec-4653-93d6-72c93aa682ba",
				"rel": "bookmark"
			}
		],
		"key_name": null,
		"image": {
			"id": "f90f6034-2570-4974-8351-6b49732ef2eb",
			"links": [
				{
					"href": "http://104.130.131.164:8774/fcad67a6189847c4aecfa3c81a05783b/images/f90f6034-2570-4974-8351-6b49732ef2eb",
					"rel": "bookmark"
				}
			]
		},
		"OS-EXT-STS:task_state": null,
		"OS-EXT-STS:vm_state": "active",
		"OS-EXT-SRV-ATTR:instance_name": "instance-0000001d",
		"OS-SRV-USG:launched_at": "2014-09-25T13:04:49.000000",
		"OS-EXT-SRV-ATTR:hypervisor_hostname": "devstack",
		"flavor": {
			"id": "1",
			"links": [
				{
					"href": "http://104.130.131.164:8774/fcad67a6189847c4aecfa3c81a05783b/flavors/1",
					"rel": "bookmark"
				}
			]
		},
		"id": "9e5476bd-a4ec-4653-93d6-72c93aa682ba",
		"security_groups": [
			{
				"name": "default"
			}
		],
		"OS-SRV-USG:terminated_at": null,
		"OS-EXT-AZ:availability_zone": "nova",
		"user_id": "9349aff8be7545ac9d2f1d00999a23cd",
		"name": "derp",
		"created": "2014-09-25T13:04:41Z",
		"tenant_id": "fcad67a6189847c4aecfa3c81a05783b",
		"OS-DCF:diskConfig": "MANUAL",
		"os-extended-volumes:volumes_attached": [],
		"accessIPv4": "",
		"accessIPv6": "",
		"progress": 0,
		"OS-EXT-STS:power_state": 1,
		"config_drive": "",
		"metadata": {},
		"fault": {
			"message": "Conflict updating instance c2ce4dea-b73f-4d01-8633-2c6032869281. Expected: {'task_state': [u'spawning']}. Actual: {'task_state': None}",
			"code": 500,
			"created": "2017-11-11T07:58:39Z",
			"details": "Stock details for test"
		}
	}
}
`

const ServerPasswordBody = `
{
    "password": "xlozO3wLCBRWAa2yDjCCVx8vwNPypxnypmRYDa/zErlQ+EzPe1S/Gz6nfmC52mOlOSCRuUOmG7kqqgejPof6M7bOezS387zjq4LSvvwp28zUknzy4YzfFGhnHAdai3TxUJ26pfQCYrq8UTzmKF2Bq8ioSEtVVzM0A96pDh8W2i7BOz6MdoiVyiev/I1K2LsuipfxSJR7Wdke4zNXJjHHP2RfYsVbZ/k9ANu+Nz4iIH8/7Cacud/pphH7EjrY6a4RZNrjQskrhKYed0YERpotyjYk1eDtRe72GrSiXteqCM4biaQ5w3ruS+AcX//PXk3uJ5kC7d67fPXaVz4WaQRYMg=="
}
`

const ConsoleOutputBody = `{
	"output": "abc"
}`

const ServerWithTagsCreateRequest = `
{
  "server": {
    "name": "derp",
    "imageRef": "f90f6034-2570-4974-8351-6b49732ef2eb",
    "flavorRef": "1",
    "tags": ["foo", "bar"]
  }
}`

// SingleServerWithTagsBody is the canned body of a Get request on an existing server with tags.
const SingleServerWithTagsBody = `
{
	"server": {
		"status": "ACTIVE",
		"updated": "2014-09-25T13:04:49Z",
		"hostId": "29d3c8c896a45aa4c34e52247875d7fefc3d94bbcc9f622b5d204362",
		"OS-EXT-SRV-ATTR:host": "devstack",
		"addresses": {
			"private": [
				{
					"OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:9e:89:be",
					"version": 4,
					"addr": "10.0.0.31",
					"OS-EXT-IPS:type": "fixed"
				}
			]
		},
		"links": [
			{
				"href": "http://104.130.131.164:8774/v2/fcad67a6189847c4aecfa3c81a05783b/servers/9e5476bd-a4ec-4653-93d6-72c93aa682ba",
				"rel": "self"
			},
			{
				"href": "http://104.130.131.164:8774/fcad67a6189847c4aecfa3c81a05783b/servers/9e5476bd-a4ec-4653-93d6-72c93aa682ba",
				"rel": "bookmark"
			}
		],
		"key_name": null,
		"image": {
			"id": "f90f6034-2570-4974-8351-6b49732ef2eb",
			"links": [
				{
					"href": "http://104.130.131.164:8774/fcad67a6189847c4aecfa3c81a05783b/images/f90f6034-2570-4974-8351-6b49732ef2eb",
					"rel": "bookmark"
				}
			]
		},
		"OS-EXT-STS:task_state": null,
		"OS-EXT-STS:vm_state": "active",
		"OS-EXT-SRV-ATTR:instance_name": "instance-0000001d",
		"OS-SRV-USG:launched_at": "2014-09-25T13:04:49.000000",
		"OS-EXT-SRV-ATTR:hypervisor_hostname": "devstack",
		"flavor": {
			"id": "1",
			"links": [
				{
					"href": "http://104.130.131.164:8774/fcad67a6189847c4aecfa3c81a05783b/flavors/1",
					"rel": "bookmark"
				}
			]
		},
		"id": "9e5476bd-a4ec-4653-93d6-72c93aa682ba",
		"security_groups": [
			{
				"name": "default"
			}
		],
		"OS-SRV-USG:terminated_at": null,
		"OS-EXT-AZ:availability_zone": "nova",
		"user_id": "9349aff8be7545ac9d2f1d00999a23cd",
		"name": "derp",
		"created": "2014-09-25T13:04:41Z",
		"tenant_id": "fcad67a6189847c4aecfa3c81a05783b",
		"OS-DCF:diskConfig": "MANUAL",
		"os-extended-volumes:volumes_attached": [],
		"accessIPv4": "",
		"accessIPv6": "",
		"progress": 0,
		"OS-EXT-STS:power_state": 1,
		"config_drive": "",
		"metadata": {},
		"tags": ["foo", "bar"]
	}
}
`

var (
	herpTimeCreated, _ = time.Parse(time.RFC3339, "2014-09-25T13:10:02Z")
	herpTimeUpdated, _ = time.Parse(time.RFC3339, "2014-09-25T13:10:10Z")
	// ServerHerp is a Server struct that should correspond to the first result in ServerListBody.
	ServerHerp = servers.Server{
		Status:  "ACTIVE",
		Updated: herpTimeUpdated,
		HostID:  "29d3c8c896a45aa4c34e52247875d7fefc3d94bbcc9f622b5d204362",
		Addresses: map[string]any{
			"private": []any{
				map[string]any{
					"OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:7c:1b:2b",
					"version":                 float64(4),
					"addr":                    "10.0.0.32",
					"OS-EXT-IPS:type":         "fixed",
				},
			},
		},
		Links: []any{
			map[string]any{
				"href": "http://104.130.131.164:8774/v2/fcad67a6189847c4aecfa3c81a05783b/servers/ef079b0c-e610-4dfb-b1aa-b49f07ac48e5",
				"rel":  "self",
			},
			map[string]any{
				"href": "http://104.130.131.164:8774/fcad67a6189847c4aecfa3c81a05783b/servers/ef079b0c-e610-4dfb-b1aa-b49f07ac48e5",
				"rel":  "bookmark",
			},
		},
		Image: map[string]any{
			"id": "f90f6034-2570-4974-8351-6b49732ef2eb",
			"links": []any{
				map[string]any{
					"href": "http://104.130.131.164:8774/fcad67a6189847c4aecfa3c81a05783b/images/f90f6034-2570-4974-8351-6b49732ef2eb",
					"rel":  "bookmark",
				},
			},
		},
		Flavor: map[string]any{
			"id": "1",
			"links": []any{
				map[string]any{
					"href": "http://104.130.131.164:8774/fcad67a6189847c4aecfa3c81a05783b/flavors/1",
					"rel":  "bookmark",
				},
			},
		},
		ID:       "ef079b0c-e610-4dfb-b1aa-b49f07ac48e5",
		UserID:   "9349aff8be7545ac9d2f1d00999a23cd",
		Name:     "herp",
		Created:  herpTimeCreated,
		TenantID: "fcad67a6189847c4aecfa3c81a05783b",
		Metadata: map[string]string{},
		AttachedVolumes: []servers.AttachedVolume{
			{
				ID: "2bdbc40f-a277-45d4-94ac-d9881c777d33",
			},
		},
		SecurityGroups: []map[string]any{
			{
				"name": "default",
			},
		},
		Host:               "devstack",
		Hostname:           nil,
		HypervisorHostname: "devstack",
		InstanceName:       "instance-0000001e",
		LaunchIndex:        nil,
		ReservationID:      nil,
		RootDeviceName:     nil,
		Userdata:           nil,
		VmState:            "active",
		PowerState:         servers.RUNNING,
		LaunchedAt:         time.Date(2014, 9, 25, 13, 10, 10, 0, time.UTC),
		TerminatedAt:       time.Time{},
		DiskConfig:         servers.Manual,
		AvailabilityZone:   "nova",
	}

	derpTimeCreated, _ = time.Parse(time.RFC3339, "2014-09-25T13:04:41Z")
	derpTimeUpdated, _ = time.Parse(time.RFC3339, "2014-09-25T13:04:49Z")
	// ServerDerp is a Server struct that should correspond to the second server in ServerListBody.
	ServerDerp = servers.Server{
		Status:  "ACTIVE",
		Updated: derpTimeUpdated,
		HostID:  "29d3c8c896a45aa4c34e52247875d7fefc3d94bbcc9f622b5d204362",
		Addresses: map[string]any{
			"private": []any{
				map[string]any{
					"OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:9e:89:be",
					"version":                 float64(4),
					"addr":                    "10.0.0.31",
					"OS-EXT-IPS:type":         "fixed",
				},
			},
		},
		Links: []any{
			map[string]any{
				"href": "http://104.130.131.164:8774/v2/fcad67a6189847c4aecfa3c81a05783b/servers/9e5476bd-a4ec-4653-93d6-72c93aa682ba",
				"rel":  "self",
			},
			map[string]any{
				"href": "http://104.130.131.164:8774/fcad67a6189847c4aecfa3c81a05783b/servers/9e5476bd-a4ec-4653-93d6-72c93aa682ba",
				"rel":  "bookmark",
			},
		},
		Image: map[string]any{
			"id": "f90f6034-2570-4974-8351-6b49732ef2eb",
			"links": []any{
				map[string]any{
					"href": "http://104.130.131.164:8774/fcad67a6189847c4aecfa3c81a05783b/images/f90f6034-2570-4974-8351-6b49732ef2eb",
					"rel":  "bookmark",
				},
			},
		},
		Flavor: map[string]any{
			"id": "1",
			"links": []any{
				map[string]any{
					"href": "http://104.130.131.164:8774/fcad67a6189847c4aecfa3c81a05783b/flavors/1",
					"rel":  "bookmark",
				},
			},
		},
		ID:              "9e5476bd-a4ec-4653-93d6-72c93aa682ba",
		UserID:          "9349aff8be7545ac9d2f1d00999a23cd",
		Name:            "derp",
		Created:         derpTimeCreated,
		TenantID:        "fcad67a6189847c4aecfa3c81a05783b",
		Metadata:        map[string]string{},
		AttachedVolumes: []servers.AttachedVolume{},
		SecurityGroups: []map[string]any{
			{
				"name": "default",
			},
		},
		Host:               "devstack",
		Hostname:           nil,
		HypervisorHostname: "devstack",
		InstanceName:       "instance-0000001d",
		LaunchIndex:        nil,
		ReservationID:      nil,
		RootDeviceName:     nil,
		Userdata:           nil,
		VmState:            "active",
		PowerState:         servers.RUNNING,
		LaunchedAt:         time.Date(2014, 9, 25, 13, 04, 49, 0, time.UTC),
		TerminatedAt:       time.Time{},
		DiskConfig:         servers.Manual,
		AvailabilityZone:   "nova",
		Locked:             func() *bool { b := true; return &b }(),
	}

	ConsoleOutput = "abc"

	merpTimeCreated, _ = time.Parse(time.RFC3339, "2014-09-25T13:04:41Z")
	merpTimeUpdated, _ = time.Parse(time.RFC3339, "2014-09-25T13:04:49Z")
	// ServerMerp is a Server struct that should correspond to the second server in ServerListBody.
	ServerMerp = servers.Server{
		Status:  "ACTIVE",
		Updated: merpTimeUpdated,
		HostID:  "29d3c8c896a45aa4c34e52247875d7fefc3d94bbcc9f622b5d204362",
		Addresses: map[string]any{
			"private": []any{
				map[string]any{
					"OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:9e:89:be",
					"version":                 float64(4),
					"addr":                    "10.0.0.31",
					"OS-EXT-IPS:type":         "fixed",
				},
			},
		},
		Links: []any{
			map[string]any{
				"href": "http://104.130.131.164:8774/v2/fcad67a6189847c4aecfa3c81a05783b/servers/9e5476bd-a4ec-4653-93d6-72c93aa682ba",
				"rel":  "self",
			},
			map[string]any{
				"href": "http://104.130.131.164:8774/fcad67a6189847c4aecfa3c81a05783b/servers/9e5476bd-a4ec-4653-93d6-72c93aa682ba",
				"rel":  "bookmark",
			},
		},
		Image: nil,
		Flavor: map[string]any{
			"id": "1",
			"links": []any{
				map[string]any{
					"href": "http://104.130.131.164:8774/fcad67a6189847c4aecfa3c81a05783b/flavors/1",
					"rel":  "bookmark",
				},
			},
		},
		ID:              "9e5476bd-a4ec-4653-93d6-72c93aa682bb",
		UserID:          "9349aff8be7545ac9d2f1d00999a23cd",
		Name:            "merp",
		Created:         merpTimeCreated,
		TenantID:        "fcad67a6189847c4aecfa3c81a05783b",
		Metadata:        map[string]string{},
		AttachedVolumes: []servers.AttachedVolume{},
		SecurityGroups: []map[string]any{
			{
				"name": "default",
			},
		},
		Host:               "devstack",
		Hostname:           nil,
		HypervisorHostname: "devstack",
		InstanceName:       "instance-0000001d",
		LaunchIndex:        nil,
		ReservationID:      nil,
		RootDeviceName:     nil,
		Userdata:           nil,
		VmState:            "active",
		PowerState:         servers.RUNNING,
		LaunchedAt:         time.Date(2014, 9, 25, 13, 04, 49, 0, time.UTC),
		TerminatedAt:       time.Time{},
		DiskConfig:         servers.Manual,
		AvailabilityZone:   "nova",
	}

	faultTimeCreated, _ = time.Parse(time.RFC3339, "2017-11-11T07:58:39Z")
	DerpFault           = servers.Fault{
		Code:    500,
		Created: faultTimeCreated,
		Details: "Stock details for test",
		Message: "Conflict updating instance c2ce4dea-b73f-4d01-8633-2c6032869281. " +
			"Expected: {'task_state': [u'spawning']}. Actual: {'task_state': None}",
	}
)

type CreateOptsWithCustomField struct {
	servers.CreateOpts
	Foo string `json:"foo,omitempty"`
}

func (opts CreateOptsWithCustomField) ToServerCreateMap() (map[string]any, error) {
	b, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	return map[string]any{"server": b}, nil
}

// HandleServerNoNetworkCreationSuccessfully sets up the test server with no
// network to respond to a server creation request with a given response.
func HandleServerNoNetworkCreationSuccessfully(t *testing.T, fakeServer th.FakeServer, response string) {
	fakeServer.Mux.HandleFunc("/servers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{
			"server": {
				"name": "derp",
				"imageRef": "f90f6034-2570-4974-8351-6b49732ef2eb",
				"flavorRef": "1",
				"networks": "none"
			}
		}`)

		w.WriteHeader(http.StatusAccepted)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, response)
	})
}

// HandleServerCreationSuccessfully sets up the test server to respond to a server creation request
// with a given response.
func HandleServerCreationSuccessfully(t *testing.T, fakeServer th.FakeServer, response string) {
	fakeServer.Mux.HandleFunc("/servers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{
			"server": {
				"name": "derp",
				"imageRef": "f90f6034-2570-4974-8351-6b49732ef2eb",
				"flavorRef": "1"
			}
		}`)

		w.WriteHeader(http.StatusAccepted)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, response)
	})

	fakeServer.Mux.HandleFunc("/images/detail", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		if err := r.ParseForm(); err != nil {
			t.Errorf("Failed to parse request form %v", err)
		}
		marker := r.Form.Get("marker")
		switch marker {
		case "":
			fmt.Fprint(w, `
				{
					"images": [
						{
							"status": "ACTIVE",
							"updated": "2014-09-23T12:54:56Z",
							"id": "f3e4a95d-1f4f-4989-97ce-f3a1fb8c04d7",
							"OS-EXT-IMG-SIZE:size": 476704768,
							"name": "F17-x86_64-cfntools",
							"created": "2014-09-23T12:54:52Z",
							"minDisk": 0,
							"progress": 100,
							"minRam": 0
						},
						{
							"status": "ACTIVE",
							"updated": "2014-09-23T12:51:43Z",
							"id": "f90f6034-2570-4974-8351-6b49732ef2eb",
							"OS-EXT-IMG-SIZE:size": 13167616,
							"name": "cirros-0.3.2-x86_64-disk",
							"created": "2014-09-23T12:51:42Z",
							"minDisk": 0,
							"progress": 100,
							"minRam": 0
						}
					]
				}
			`)
		case "2":
			fmt.Fprint(w, `{ "images": [] }`)
		default:
			t.Fatalf("Unexpected marker: [%s]", marker)
		}
	})

	fakeServer.Mux.HandleFunc("/flavors/detail", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		if err := r.ParseForm(); err != nil {
			t.Errorf("Failed to parse request form %v", err)
		}
		marker := r.Form.Get("marker")
		switch marker {
		case "":
			fmt.Fprintf(w, `
						{
							"flavors": [
								{
									"id": "1",
									"name": "m1.tiny",
									"disk": 1,
									"ram": 512,
									"vcpus": 1,
									"swap":""
								},
								{
									"id": "2",
									"name": "m2.small",
									"disk": 10,
									"ram": 1024,
									"vcpus": 2,
									"swap": 1000
								}
							],
							"flavors_links": [
								{
									"href": "%s/flavors/detail?marker=2",
									"rel": "next"
								}
							]
						}
					`, fakeServer.Server.URL)
		case "2":
			fmt.Fprint(w, `{ "flavors": [] }`)
		default:
			t.Fatalf("Unexpected marker: [%s]", marker)
		}
	})
}

// HandleServerCreationWithCustomFieldSuccessfully sets up the test server to respond to a server creation request
// with a given response.
func HandleServersCreationSuccessfully(t *testing.T, fakeServer th.FakeServer, response string) {
	fakeServer.Mux.HandleFunc("/servers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{
			"server": {
				"name": "derp",
				"imageRef": "f90f6034-2570-4974-8351-6b49732ef2eb",
				"flavorRef": "1",
				"min_count": 3,
				"max_count": 3
			}
		}`)

		w.WriteHeader(http.StatusAccepted)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, response)
	})
}

// HandleServerCreationWithCustomFieldSuccessfully sets up the test server to respond to a server creation request
// with a given response.
func HandleServerCreationWithCustomFieldSuccessfully(t *testing.T, fakeServer th.FakeServer, response string) {
	fakeServer.Mux.HandleFunc("/servers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{
			"server": {
				"name": "derp",
				"imageRef": "f90f6034-2570-4974-8351-6b49732ef2eb",
				"flavorRef": "1",
				"foo": "bar"
			}
		}`)

		w.WriteHeader(http.StatusAccepted)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, response)
	})
}

func HandleServerCreationWithHostname(t *testing.T, fakeServer th.FakeServer, response string) {
	fakeServer.Mux.HandleFunc("/servers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{
			"server": {
				"name": "derp",
				"hostname": "derp.local",
				"imageRef": "f90f6034-2570-4974-8351-6b49732ef2eb",
				"flavorRef": "1"
			}
		}`)

		w.WriteHeader(http.StatusAccepted)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, response)
	})
}

// HandleServerCreationWithUserdata sets up the test server to respond to a server creation request
// with a given response.
func HandleServerCreationWithUserdata(t *testing.T, fakeServer th.FakeServer, response string) {
	fakeServer.Mux.HandleFunc("/servers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{
			"server": {
				"name": "derp",
				"imageRef": "f90f6034-2570-4974-8351-6b49732ef2eb",
				"flavorRef": "1",
				"user_data": "dXNlcmRhdGEgc3RyaW5n"
			}
		}`)

		w.WriteHeader(http.StatusAccepted)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, response)
	})
}

// HandleServerCreationWithMetadata sets up the test server to respond to a server creation request
// with a given response.
func HandleServerCreationWithMetadata(t *testing.T, fakeServer th.FakeServer, response string) {
	fakeServer.Mux.HandleFunc("/servers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{
			"server": {
				"name": "derp",
				"imageRef": "f90f6034-2570-4974-8351-6b49732ef2eb",
				"flavorRef": "1",
				"metadata": {
					"abc": "def"
				}
			}
		}`)

		w.WriteHeader(http.StatusAccepted)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, response)
	})
}

// HandleServerListSimpleSuccessfully sets up the test server to respond to a server List request.
func HandleServerListSimpleSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	fakeServer.Mux.HandleFunc("/servers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		if err := r.ParseForm(); err != nil {
			t.Errorf("Failed to parse request form %v", err)
		}
		marker := r.Form.Get("marker")
		switch marker {
		case "":
			fmt.Fprint(w, ServerListBody)
		case "9e5476bd-a4ec-4653-93d6-72c93aa682ba":
			fmt.Fprint(w, `{ "servers": [] }`)
		default:
			t.Fatalf("/servers invoked with unexpected marker=[%s]", marker)
		}
	})
}

// HandleServerListSuccessfully sets up the test server to respond to a server detail List request.
func HandleServerListSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	fakeServer.Mux.HandleFunc("/servers/detail", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		if err := r.ParseForm(); err != nil {
			t.Errorf("Failed to parse request form %v", err)
		}
		marker := r.Form.Get("marker")
		switch marker {
		case "":
			fmt.Fprint(w, ServerListBody)
		case "9e5476bd-a4ec-4653-93d6-72c93aa682ba":
			fmt.Fprint(w, `{ "servers": [] }`)
		default:
			t.Fatalf("/servers/detail invoked with unexpected marker=[%s]", marker)
		}
	})
}

// HandleServerDeletionSuccessfully sets up the test server to respond to a server deletion request.
func HandleServerDeletionSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	fakeServer.Mux.HandleFunc("/servers/asdfasdfasdf", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})
}

// HandleServerForceDeletionSuccessfully sets up the test server to respond to a server force deletion
// request.
func HandleServerForceDeletionSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	fakeServer.Mux.HandleFunc("/servers/asdfasdfasdf/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{ "forceDelete": "" }`)

		w.WriteHeader(http.StatusAccepted)
	})
}

// HandleServerGetSuccessfully sets up the test server to respond to a server Get request.
func HandleServerGetSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	fakeServer.Mux.HandleFunc("/servers/1234asdf", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		fmt.Fprint(w, SingleServerBody)
	})
}

// HandleServerGetFaultSuccessfully sets up the test server to respond to a server Get
// request which contains a fault.
func HandleServerGetFaultSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	fakeServer.Mux.HandleFunc("/servers/1234asdf", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		fmt.Fprint(w, FaultyServerBody)
	})
}

// HandleServerUpdateSuccessfully sets up the test server to respond to a server Update request.
func HandleServerUpdateSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	fakeServer.Mux.HandleFunc("/servers/1234asdf", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestJSONRequest(t, r, `{ "server": { "name": "new-name" } }`)

		fmt.Fprint(w, SingleServerBody)
	})
}

// HandleAdminPasswordChangeSuccessfully sets up the test server to respond to a server password
// change request.
func HandleAdminPasswordChangeSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	fakeServer.Mux.HandleFunc("/servers/1234asdf/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{ "changePassword": { "adminPass": "new-password" } }`)

		w.WriteHeader(http.StatusAccepted)
	})
}

// HandleRebootSuccessfully sets up the test server to respond to a reboot request with success.
func HandleRebootSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	fakeServer.Mux.HandleFunc("/servers/1234asdf/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{ "reboot": { "type": "SOFT" } }`)

		w.WriteHeader(http.StatusAccepted)
	})
}

// HandleShowConsoleOutputSuccessfully sets up the test server to respond to a os-getConsoleOutput request with success.
func HandleShowConsoleOutputSuccessfully(t *testing.T, fakeServer th.FakeServer, response string) {
	fakeServer.Mux.HandleFunc("/servers/1234asdf/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{ "os-getConsoleOutput": { "length": 50 } }`)

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, response)
	})
}

// HandleRebuildSuccessfully sets up the test server to respond to a rebuild request with success.
func HandleRebuildSuccessfully(t *testing.T, fakeServer th.FakeServer, response string) {
	fakeServer.Mux.HandleFunc("/servers/1234asdf/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `
			{
				"rebuild": {
					"name": "new-name",
					"adminPass": "swordfish",
					"imageRef": "f90f6034-2570-4974-8351-6b49732ef2eb",
					"accessIPv4": "1.2.3.4"
				}
			}
		`)

		w.WriteHeader(http.StatusAccepted)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, response)
	})
}

// HandleMetadatumGetSuccessfully sets up the test server to respond to a metadatum Get request.
func HandleMetadatumGetSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	fakeServer.Mux.HandleFunc("/servers/1234asdf/metadata/foo", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		_, err := w.Write([]byte(`{ "meta": {"foo":"bar"}}`))
		th.AssertNoErr(t, err)
	})
}

// HandleMetadatumCreateSuccessfully sets up the test server to respond to a metadatum Create request.
func HandleMetadatumCreateSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	fakeServer.Mux.HandleFunc("/servers/1234asdf/metadata/foo", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{
			"meta": {
				"foo": "bar"
			}
		}`)

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		_, err := w.Write([]byte(`{ "meta": {"foo":"bar"}}`))
		th.AssertNoErr(t, err)
	})
}

// HandleMetadatumDeleteSuccessfully sets up the test server to respond to a metadatum Delete request.
func HandleMetadatumDeleteSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	fakeServer.Mux.HandleFunc("/servers/1234asdf/metadata/foo", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})
}

// HandleMetadataGetSuccessfully sets up the test server to respond to a metadata Get request.
func HandleMetadataGetSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	fakeServer.Mux.HandleFunc("/servers/1234asdf/metadata", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{ "metadata": {"foo":"bar", "this":"that"}}`))
		th.AssertNoErr(t, err)
	})
}

// HandleMetadataResetSuccessfully sets up the test server to respond to a metadata Create request.
func HandleMetadataResetSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	fakeServer.Mux.HandleFunc("/servers/1234asdf/metadata", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{
				"metadata": {
					"foo": "bar",
					"this": "that"
				}
			}`)

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		_, err := w.Write([]byte(`{ "metadata": {"foo":"bar", "this":"that"}}`))
		th.AssertNoErr(t, err)
	})
}

// HandleMetadataUpdateSuccessfully sets up the test server to respond to a metadata Update request.
func HandleMetadataUpdateSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	fakeServer.Mux.HandleFunc("/servers/1234asdf/metadata", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{
				"metadata": {
					"foo": "baz",
					"this": "those"
				}
			}`)

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		_, err := w.Write([]byte(`{ "metadata": {"foo":"baz", "this":"those"}}`))
		th.AssertNoErr(t, err)
	})
}

// ListAddressesExpected represents an expected repsonse from a ListAddresses request.
var ListAddressesExpected = map[string][]servers.Address{
	"public": {
		{
			Version: 4,
			Address: "50.56.176.35",
		},
		{
			Version: 6,
			Address: "2001:4800:790e:510:be76:4eff:fe04:84a8",
		},
	},
	"private": {
		{
			Version: 4,
			Address: "10.180.3.155",
		},
	},
}

// HandleAddressListSuccessfully sets up the test server to respond to a ListAddresses request.
func HandleAddressListSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	fakeServer.Mux.HandleFunc("/servers/asdfasdfasdf/ips", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, `{
			"addresses": {
				"public": [
				{
					"version": 4,
					"addr": "50.56.176.35"
				},
				{
					"version": 6,
					"addr": "2001:4800:790e:510:be76:4eff:fe04:84a8"
				}
				],
				"private": [
				{
					"version": 4,
					"addr": "10.180.3.155"
				}
				]
			}
		}`)
	})
}

// ListNetworkAddressesExpected represents an expected repsonse from a ListAddressesByNetwork request.
var ListNetworkAddressesExpected = []servers.Address{
	{
		Version: 4,
		Address: "50.56.176.35",
	},
	{
		Version: 6,
		Address: "2001:4800:790e:510:be76:4eff:fe04:84a8",
	},
}

// HandleNetworkAddressListSuccessfully sets up the test server to respond to a ListAddressesByNetwork request.
func HandleNetworkAddressListSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	fakeServer.Mux.HandleFunc("/servers/asdfasdfasdf/ips/public", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, `{
			"public": [
			{
				"version": 4,
				"addr": "50.56.176.35"
				},
				{
					"version": 6,
					"addr": "2001:4800:790e:510:be76:4eff:fe04:84a8"
				}
			]
			}`)
	})
}

// HandleCreateServerImageSuccessfully sets up the test server to respond to a TestCreateServerImage request.
func HandleCreateServerImageSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	fakeServer.Mux.HandleFunc("/servers/serverimage/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		w.Header().Add("Location", "https://0.0.0.0/images/xxxx-xxxxx-xxxxx-xxxx")
		w.WriteHeader(http.StatusAccepted)
	})
}

// HandlePasswordGetSuccessfully sets up the test server to respond to a password Get request.
func HandlePasswordGetSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	fakeServer.Mux.HandleFunc("/servers/1234asdf/os-server-password", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")

		fmt.Fprint(w, ServerPasswordBody)
	})
}

// HandleServerWithTagsCreationSuccessfully sets up the test server to respond
// to a server creation request with a given response.
func HandleServerWithTagsCreationSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	fakeServer.Mux.HandleFunc("/servers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, ServerWithTagsCreateRequest)

		w.WriteHeader(http.StatusAccepted)
		w.Header().Add("Content-Type", "application/json")

		fmt.Fprint(w, SingleServerWithTagsBody)
	})
}

// HandleServerHostnameUpdateSuccessfully sets up the test server to respond to a server update
// request changing the hostname.
func HandleServerHostnameUpdateSuccessfully(t *testing.T, fakeServer th.FakeServer) {
	fakeServer.Mux.HandleFunc("/servers/1234asdf", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestJSONRequest(t, r, `{ "server": { "hostname": "new-hostname" } }`)

		fmt.Fprint(w, SingleServerBody)
	})
}
