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
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

As a backend/ml engineer, I have worked with many kinds of databases. SQL/NoSQL, KeyValue, Document, Read-Heavy, and Write-Heavy databases. However, I never had a chance to really understand the implementation details of the databases. Like indexing, search, caching, versioning, and even distributed data store. Therefore, I decided to implement my own databases, start from the most simplest one, the Bitcask database and keep involving to distributed K/V store. 

### Design Spec:
Briefly describe the requirements and secnerio.
* Key Value storage
* Write-heavy
* The size of data should be 10 times larger than memory.
* Support crash recovery
* Distributed storage

<p align="right">(<a href="#readme-top">back to top</a>)</p>


### Built With:
This section lists the the main packages, In general most of the core functions were built from the ground.
In other words, re-inventing the wheel lol.  Such as SSTables, Segment file, Primary index ...etc.

* [![Gin][gin-gonic]][gin-url]
* [![gRPC-go][gRPC]][gRPC-url]
* [![Gin-Swagger][swagger]][swagger-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running follow these simple example steps.

### Prerequisites

This is an example of how to list things you need to use the software and how to install them.
* npm
  ```sh
  npm install npm@latest -g
  ```

### Installation

_Below is an example of how you can instruct your audience on installing and setting up your app. This template doesn't rely on any external dependencies or services._

1. Get a free API Key at [https://example.com](https://example.com)
2. Clone the repo
   ```sh
   git clone https://github.com/your_username_/Project-Name.git
   ```
3. Install NPM packages
   ```sh
   npm install
   ```
4. Enter your API in `config.js`
   ```js
   const API_KEY = 'ENTER YOUR API';
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

Use this space to show useful examples of how a project can be used. Additional screenshots, code examples and demos work well in this space. You may also link to more resources.

_For more examples, please refer to the [Documentation](https://example.com)_

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

### Implementation Roadmap
- [x]  Basic Get / Set / Delete Methods 
- [x]  Implement vanilla hashtable key value storage
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

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Contact
[@Email](lochuhsin@gmail.com)
[@linkedin](https://www.linkedin.com/in/lochuhsin/)

<p align="right">(<a href="#readme-top">back to top</a>)</p>




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
