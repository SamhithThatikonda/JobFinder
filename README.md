# GOLANG GRAPHQL MONGODB CRUD Project

This is the accompanying code for my YouTube video with the same name (almost the same name).

## Project Initialization

### Step 1: Create a New Folder for the Project
```bash
mkdir gql-yt
```
### Step 2: Initialize Your Project
```bash
cd gql-yt
go mod init github.com/akhil/gql-yt
```

### Step 3: Get gqlgen for Your Project
```bash
go get github.com/99designs/gqlgen
```

### Step 4: Add gqlgen to tools.go
```bash
printf '// +build tools\npackage tools\nimport _ "github.com/99designs/gqlgen"' | gofmt > tools.go
```

### Step 5: Get All the Dependencies
```bash
go mod tidy
```
### Step 6: Initialize Your Project and Generate Code After Writing the GraphQL Schema
```bash
go run github.com/99designs/gqlgen init
go run github.com/99designs/gqlgen generate
```

GraphQL Queries to Interact with the API

Get All Jobs

```graphql
query GetAllJobs {
  jobs {
    _id
    title
    description
    company
    url
  }
}
```

Create Job

```graphql
mutation CreateJobListing($input: CreateJobListingInput!) {
  createJobListing(input: $input) {
    _id
    title
    description
    company
    url
  }
}
```

Example Input:
```json
{
  "input": {
    "title": "Software Development Engineer - I",
    "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt",
    "company": "Google",
    "url": "www.google.com/"
  }
}
```

Get Job By Id
```graphql
query GetJob($id: ID!) {
  job(id: $id) {
    _id
    title
    description
    url
    company
  }
}
```

Example Input:
```json
{
  "id": "638051d7acc418c13197fdf7"
}
```

Update Job By Id
```graphql
mutation UpdateJobListing($id: ID!, $input: UpdateJobListingInput!) {
  updateJobListing(id: $id, input: $input) {
    title
    description
    _id
    company
    url
  }
}
```

Example Input:
```json
{
  "id": "638051d3acc418c13197fdf6",
  "input": {
    "title": "Software Development Engineer - III"
  }
}
```

Delete Job By Id
```graphql
mutation DeleteJobListing($id: ID!) {
  deleteJobListing(id: $id) {
    deletedJobId
  }
}
```

Example Input:
```json
{
  "id": "638051d3acc418c13197fdf6"
}
```

This README file provides step-by-step instructions to initialize the project and the GraphQL queries to interact with the API.
