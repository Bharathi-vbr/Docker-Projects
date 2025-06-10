# 🌐 Docker Networking

Docker networking allows containers to communicate with each other, the host system, and external networks. Since containers are isolated by default, Docker provides networking capabilities to enable secure and flexible communication.

---

## 🚀 Default Docker Networks

You can list existing Docker networks using:

```
docker network ls

```
Typical output:

NETWORK ID          NAME                DRIVER
xxxxxxxxxxxx        none                null
xxxxxxxxxxxx        host                host
xxxxxxxxxxxx        bridge              bridge

### 🔗 Bridge Networking (Default)
The bridge network is Docker’s default network driver for standalone containers. It provides a private internal network on the host, allowing containers on the same bridge to communicate with each other.

🧠 By default, all containers are connected to the bridge network unless specified otherwise.

### 📸 Network Topology
![image](https://github.com/user-attachments/assets/79e5e695-a892-46ba-976a-df5ab393e8b5)


### 🧱 Create a Custom Bridge Network
You can isolate containers from the default bridge network by creating your own:


```
docker network create -d bridge my_bridge

```
### Verify:


```
docker network ls

```
### You’ll see:

```
NETWORK ID          NAME                DRIVER
xxxxxxxxxxxx        bridge              bridge
xxxxxxxxxxxx        my_bridge           bridge
xxxxxxxxxxxx        none                null
xxxxxxxxxxxx        host                host

```
### 🧩 Run Containers on a Custom Network

```
docker run -d --net=my_bridge --name db training/postgres

```
This attaches the container to my_bridge. Containers on different networks are isolated and cannot communicate with each other.

### To enable communication between isolated containers:

```
docker network connect my_bridge web

```
### ✅ Now both web and db can talk over my_bridge.

### 🖥️ Host Networking
The host network mode removes network isolation. Containers share the host’s networking stack (IP address, ports, etc).

➕ Run a Container with Host Networking

```
docker run --network="host" <image_name> <command>

```
### ⚠️ Security Warning
Containers using the host network have full access to the host’s network stack.

This increases performance for some applications, but also reduces security and isolation.

### ℹ️ Not all images support host networking. Test with --network="bridge" first.

### 🌉 Overlay Networking
Designed for multi-host Docker setups (e.g., Docker Swarm).

Allows containers running on different hosts to communicate over a single virtual network.

Use Case:

Connecting services in a microservice architecture across a distributed system.

### 🧭 Macvlan Networking
Assigns a MAC address and IP to containers, making them appear as physical devices on the network.

Useful when your application requires containers to be directly reachable on the physical network.

Use Case:

IoT gateways, legacy systems integration, or network policies that require physical-like devices.

### 📚 Summary of Network Drivers
| Network Driver | Scope       | Use Case                                 | Isolation | Notes                                    |
| -------------- | ----------- | ---------------------------------------- | --------- | ---------------------------------------- |
| `bridge`       | Single host | Default; local container communication   | High      | Simple and effective                     |
| `host`         | Single host | High-performance local networking        | Low       | No isolation, use with caution           |
| `overlay`      | Multi-host  | Swarm mode, multi-host container linking | Medium    | Requires Docker Swarm or orchestration   |
| `macvlan`      | Single host | Container appears as physical host       | Medium    | Advanced use, requires MAC/IP assignment |
| `none`         | Single host | Disable networking                       | Full      | Useful for locked-down containers        |

### ✅ Best Practices
Use bridge networks for local development and testing.

Use overlay networks in distributed or orchestrated environments (e.g., Docker Swarm).

Use host networking only when performance is critical and security isn’t a concern.

Use macvlan for network-level control and when containers need direct access to the LAN.

