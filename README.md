<!-- PROJECT LOGO -->
<br />
<div align="center">
    <img src="" alt="Logo" width="80" height="80">

  <h3 align="center">Simple Job Site</h3>

  <p align="center">
    An awesome README template to jumpstart your projects!
    <br />
    <a href="https://github.com/othneildrew/Best-README-Template"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/othneildrew/Best-README-Template">View Demo</a>
    ·
    <a href="https://github.com/othneildrew/Best-README-Template/issues/new?labels=bug&template=bug-report---.md">Report Bug</a>
    ·
    <a href="https://github.com/othneildrew/Best-README-Template/issues/new?labels=enhancement&template=feature-request---.md">Request Feature</a>
  </p>
</div>
# vignesh-aug-2024

# Job Site


REST API for JobSite using golang, gin framework & GORM ORM.


## Overview
- There are 2 different roles, User and admin
- A user can apply posts, view their job applied details ,get all job posts (User can see all job updates,view their own details)
- A admin creates the JobPost,update the job posts,view the applicant details(Admin can update their own posts)



## Features
- User authentication and authorization using JSON Web Tokens (JWT)
- Admins can only see their company details and view their company applied users only
- Users can get the all company post and they can able to apply for this jobs
- Error handling and response formatting
- Input validation and data sanitization
- Database integration using PostgreSQL


## Requirements
- Golang 
- Postgres


## Run Locally

Clone the project

```bash
  git clone https://github.com/marees7/vignesh-aug-2024.git
```

Go to the project directory
go to the cmd folder and main.go file.
change the credentials of postgres db in the internals.

```bash
  go run main.go
```


## API Endpoints

The following endpoints are available in the API:

## AUTH API

| Method | 	Endpoint | 	Description |
| ---- | -------- | -------- |
| POST |	/auth/signup	| Register a new user |
| POST |	/auth/login	| Log in and obtain JWT |


## ADMIN API

| Method | 	Endpoint | 	Description |
| ---- | -------- | -------- |
| POST |	/admin/insert/:admin_id | Create a new posts |
| PUT  |	/admin/userjobsbyid/:job_id/:admin_id	| Update thier own posts | 
| GET  |	/admin/userid/:user_id/:admin_id |Get Jobs By Admin for who applied in their jobposts|
| GET   |   /admin/userdetails/:job_role/:admin_id  | Get by role and Userid|
| GET   |   /admin/postdetails/:admin_id    |Get Jobs for their Own AdminIds|
| GET   |  /admin/userjobsbyid/:job_id/:admin_id   | Get by JobId and UserId|

## USER API

| Method | 	Endpoint | 	Description |
| ---- | -------- | -------- |
| GET  |	/users/allposts	| user or admin Get all job post details |
| GET  |	/users/jobs/:job_title/:country	| Get the jobs by their country |
| GET  |	/users/company/:company_name	| Get all posts in the company |
| POST |	/users/post/:user_id    | Apply JobPost for the Job |
| GET  |	/usersowndetails/user/:user_id| user get by their userowndetails |



## Database Schema

The application uses a PostgreSQL database with the following schema:

```sql
CREATE TABLE IF NOT EXISTS public.user_details
(
    user_id bigint NOT NULL DEFAULT nextval('user_details_user_id_seq'::regclass),
    name character varying(100) COLLATE pg_catalog."default",
    email character varying(100) COLLATE pg_catalog."default",
    password character varying(255) COLLATE pg_catalog."default",
    phone_number character varying(15) COLLATE pg_catalog."default",
    role_type character varying(25) COLLATE pg_catalog."default",
    token character varying(255) COLLATE pg_catalog."default",
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    CONSTRAINT user_details_pkey PRIMARY KEY (user_id),
    CONSTRAINT uni_user_details_email UNIQUE (email)
)
CREATE TABLE IF NOT EXISTS public.job_creations
(
    job_id bigint NOT NULL DEFAULT nextval('job_creations_job_id_seq'::regclass),
    domain_id bigint,
    company_name character varying(100) COLLATE pg_catalog."default",
    company_email character varying(100) COLLATE pg_catalog."default",
    job_title character varying(100) COLLATE pg_catalog."default",
    job_status character varying(100) COLLATE pg_catalog."default",
    job_time character varying(50) COLLATE pg_catalog."default",
    description text COLLATE pg_catalog."default",
    experience bigint,
    skills character varying(255) COLLATE pg_catalog."default",
    vacancy bigint,
    "country " character varying(20) COLLATE pg_catalog."default",
    street character varying(255) COLLATE pg_catalog."default",
    city character varying(100) COLLATE pg_catalog."default",
    state character varying(100) COLLATE pg_catalog."default",
    zip_code character varying(20) COLLATE pg_catalog."default",
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    CONSTRAINT job_creations_pkey PRIMARY KEY (job_id)
)

CREATE TABLE IF NOT EXISTS public.user_job_details
(
    user_id bigint,
    job_id bigint,
    experience bigint,
    skills character varying(255) COLLATE pg_catalog."default",
    language character varying(255) COLLATE pg_catalog."default",
    country character varying(255) COLLATE pg_catalog."default",
    job_role character varying(255) COLLATE pg_catalog."default",
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    CONSTRAINT fk_user_job_details_job FOREIGN KEY (job_id)
        REFERENCES public.job_creations (job_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT fk_user_job_details_user FOREIGN KEY (user_id)
        REFERENCES public.user_details (user_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)
```

## Dependencies

The project utilizes the following third-party libraries:

- `gin-gonic/gin`: HTTP web framework
- `joho/godotenv`: Environment variable loading
- `golang-jwt/jwt/v5`: JWT implementation
- `gorm.io/gorm`: PostgreSQL ORM
- `gorm.io/driver/postgres`: PostgreSQL extensions for gorm

Make sure to run go mod download as mentioned in the installation steps to fetch these dependencies.


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
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

[![Product Name Screen Shot][product-screenshot]](https://example.com)

There are many great README templates available on GitHub; however, I didn't find one that really suited my needs so I created this enhanced one. I want to create a README template so amazing that it'll be the last one you ever need -- I think this is it.

Here's why:
* Your time should be focused on creating something amazing. A project that solves a problem and helps others
* You shouldn't be doing the same tasks over and over like creating a README from scratch
* You should implement DRY principles to the rest of your life :smile:

Of course, no one template will serve all projects since your needs may be different. So I'll be adding more in the near future. You may also suggest changes by forking this repo and creating a pull request or opening an issue. Thanks to all the people have contributed to expanding this template!

Use the `BLANK_README.md` to get started.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

This section should list any major frameworks/libraries used to bootstrap your project. Leave any add-ons/plugins for the acknowledgements section. Here are a few examples.

* [![Next][Next.js]][Next-url]
* [![React][React.js]][React-url]
* [![Vue][Vue.js]][Vue-url]
* [![Angular][Angular.io]][Angular-url]
* [![Svelte][Svelte.dev]][Svelte-url]
* [![Laravel][Laravel.com]][Laravel-url]
* [![Bootstrap][Bootstrap.com]][Bootstrap-url]
* [![JQuery][JQuery.com]][JQuery-url]

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
   git clone https://github.com/github_username/repo_name.git
   ```
3. Install NPM packages
   ```sh
   npm install
   ```
4. Enter your API in `config.js`
   ```js
   const API_KEY = 'ENTER YOUR API';
   ```
5. Change git remote url to avoid accidental pushes to base project
   ```sh
   git remote set-url origin github_username/repo_name
   git remote -v # confirm the changes
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->
## Contact

Your Name - [/Vigneshwartt) -

Project Link: [https://github.com/your_username/repo_name](https://github.com/your_username/repo_name)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



