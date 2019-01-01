# Spill

CLI for syncing secure notes between 1Password and your local filesystem.

## Use Case

Sharing env config for a project with other developers on the same project via
a shared 1Password vault. Typically syncing config is manual and therefore
error prone. There are two scenarios to avoid:

1. Making changes to the config for an environment and forgetting to update
   1Password.

2. Performing a build and forgetting to pull down the latest config from
   1Password when another developer has made a change.

## Usage

Pulling from 1Password:

```sh
spill pull .env.staging
```

Pushing to 1Password:

```sh
spill push .env.staging
```

In both cases if there are differences between the local version and the
version in 1Password a diff will be displayed and confirmation requested before
syncing.

## Running Locally

```sh
go run spill.go <subcommand>
```
