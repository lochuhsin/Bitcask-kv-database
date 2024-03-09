<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/othneildrew/Best-README-Template">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">Re Bitcask</h3>

  <p align="center">
    The journey to re implement bitcask and turn into distributed kv database.
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
        <li><a href="#run">Run</a></li>
        <li><a href="#usage">Usage</a></li>
      </ul>
    </li>
    <li><a href="#design">Design Overview</a></li>
      <ul>
        <li><a href="#overview">Overview</a></li>
        <li><a href="#components">Components</a></li>
      </ul>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

As a Backend/ML engineer, I have worked with many kinds of databases. SQL/NoSQL, KeyValue, Document, Read-Heavy, and Write-Heavy databases. 

However, I never had a chance to really understand the implementation details of the databases. Like indexing, search, caching, versioning, and even distributed data store. 

Therefore, I decided to implement my own databases, start from the most simplest one, the BitCask database and keep involving to distributed K/V store.

Of course, this project is educational purpose, still I'm trying my best to make it more robust.

### Design Spec:
This section briefly describe the requirements and scenarios.
* Key Value storage
* Write-heavy
* The size of data should be 10 times larger than memory.
* Support crash recovery
* Distributed storage


### Built With:
This section lists the the main packages, In general most of the core functions were built from the ground.
In other words, re-inventing the wheel lol.  Such as SSTables, Segment file, Primary index ...etc.

* [![Gin][gin-gonic]][gin-url]
* [![gRPC-go][gRPC]][gRPC-url]
* [![Gin-Swagger][swagger]][swagger-url]



<!-- GETTING STARTED -->
## Getting Started

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running follow these simple example steps.

### Prerequisites
- go environment

- Using as web server
  - `docker`

- Using as cluster
  - `docker`

### Run
- Standalone container web server run `make compose-up`
- Local web server run `make init make build && make run`
- Cluster setup `make cluster-up`
- Embed in project, see `endpoints.go` for public api


<!-- USAGE EXAMPLES -->
### Usage
```go
TBD
```


<!-- Design -->
## Design
### Overview
  ![alt text](./images/Database%20Overview.png)
### Components
  - Memory Manager
  - Segment Manager



<!-- ROADMAP -->
## Roadmap

### Implementation Roadmap
- [x]  Basic Get / Set / Delete Methods 
- [x]  Implement vanilla hash table key value storage
- [x]  Implement Segment storage,
- [x]  Implement SSTable (**Sorted String Table**)
- [x]  Implement Binary Search Tree for in memory storage
- [x]  Implement BloomFilter and cache mechanism for Read (Drop currently, might move to segment level)
- [x]  Implement Asynchronous TaskPool for SSTable
- [x]  Implement Primary Index for Segments
- [x]  Setup gRPC for cluster storage communication
- [ ]  Implement Red-Black Tree for in memory storage
- [ ]  Implement Raft algorithm
- [ ]  Implement Graceful exit, Crash recovery
- [ ]  Implement Range based key query
- [ ]  Add more tests for each package.



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.



<!-- CONTACT -->
## Contact
- [@Email](lochuhsin@gmail.com)
- [@linkedin](https://www.linkedin.com/in/lochuhsin/)




<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[license-shield]: https://img.shields.io/github/license/othneildrew/Best-README-Template.svg?style=for-the-badge
[license-url]: https://github.com/othneildrew/Best-README-Template/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/othneildrew
[gin-gonic]:https://img.shields.io/badge/gin-gonic?style=for-the-badge&logo=gin&logoColor=trasnparent&labelColor=black&color=black
[gin-url]: https://github.com/gin-gonic/gin?tab=readme-ov-file
[gRPC]: https://img.shields.io/badge/grpc-go?style=for-the-badge&logo=grpc-go&logoColor=trasnparent&labelColor=black&color=black
[gRPC-url]:https://github.com/grpc/grpc-go
[React.js]: https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB
[swagger]: https://img.shields.io/badge/gin-swagger?style=for-the-badge&logo=swagger&logoColor=trasnparent&labelColor=black&color=black
[swagger-url]:https://github.com/swaggo/gin-swagger
