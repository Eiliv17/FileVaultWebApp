<a name="readme-top"></a>
<!-- Head section -->
<div align="center">
  <h1 align="center">File Vault Web App</h1>

  <p align="center">
    A simple web app to password lock your files.
  </p>
</div>


<!-- Table of contents -->
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
        <li><a href="#mongodb-server">MongoDB Server</a></li>
        <li><a href="#env-file">.env File</a></li>
      </ul>
    </li>
  </ol>
</details>

<!-- About the project section -->
## About The Project

This is a simple project that allows you to upload a file on the web server and use a password to protect it, later you can download your file if you chose the correct password.

This project was build with the Gin framework and the Go programming languages. This still needs a design rework with support for mobile devices and maybe some code refactoring but as of now the app works.

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- Built with section -->
### Built With

This project was built with the following technologies:

- ![HTML](https://img.shields.io/badge/HTML5-E34F26?style=for-the-badge&logo=html5&logoColor=white)
- ![CSS](https://img.shields.io/badge/CSS3-1572B6?style=for-the-badge&logo=css3&logoColor=white)
- ![Golang](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
- ![MongoDB](https://img.shields.io/badge/MongoDB-4EA94B?style=for-the-badge&logo=mongodb&logoColor=white)

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- Getting started section -->
## Getting Started
Before you run the app you need to have the following:
- MongoDB Server
- .env File

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### MongoDB Server
You need to have MongoDB installed on your machine, if you don't have it installed you can download the installer at the following [link](https://www.mongodb.com/try/download/community). 

This link is for the community version of MongoDB, once you've completed the installation process, it should be already up and running.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### .env File
In the root of your project folder you also need to create a .env file which will load the environment variables that are needed for the web app to work. 

The app requires the following environment variables:

Variable                    | Description
---                         | ---
PORT                        | The HTTP server port
MONGODB_URI                 | The URI for connecting to your MongoDB server
DB_NAME                     | The name of the database (inside MongoDB) where you want to store the files
FILE_DIR                    | The name of the directory where the uploaded files will be stored

Here's an example of the .env file:
```
PORT=3000
MONGODB_URI= mongodb://127.0.0.1:27017/
DB_NAME=VaultWebApp
FILE_DIR=filedb
```
<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- Usage section -->
## Usage
If you've followed the getting started correctly, you'll then be able to run the application by just compiling and running the main.go file.

<p align="right">(<a href="#readme-top">back to top</a>)</p>