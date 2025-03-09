# GitHub Organization Bootstrap

## Pre-Requisites

- A GitHub organization to bootstrap
- An authenticated GitHub client with appropriate permissions. Some scopes (organization ones) are _unusual_ and require to be specifically requested when authenticating with GitHub CLI. You can use the following command to authenticate with the required scopes:

```sh
gh auth login --scopes 'repo workflow admin:org'
```

## Missing IaC

Some parts of the GitHub organization are not managed by IaC, as GitHub does not provide a way to manage them through Pulumi nor Terraform.

- General
  - GitHub Developer Program > Enroll
  - GitHub Sponsors > Set up
- Organization roles
  - Role assignments
    - `admins` > `All-repository admin`
    - `maintainers` > `All-repository maintain` & `Security manager`
    - `developers` > `All-repository read`
- Member privileges
  - Repository discussions > `Enable`
  - Projects base permissions > `Write`
  - Integration access requests
    - Allow integration requests from outside collaborators > `Disable`
  - Admin repository permissions
    - Repository visibility change
      - Allow members to change repository visibilities for this organization > `Disable`
    - Repository deletion and transfer
      - Allow members to delete or transfer repositories for this organization > `Disable`
    - Issue deletion
      - Allow members to delete issues for this organization > `Disable`
  - Member team permissions
    - Team creation rules
    - Allow members to create teams > `Disable`
- Repository
  - Repository labels > Delete all
- Codespaces
  - Codespaces access > `Enable for specific members or teams`
- Planning
  - Projects
    - Allow members to change project visibilities for this organization > `Disable`
  - Issue types
    - `Bug`: `Unexpected problem or behavior` - Orange
    - `Feature`: `New functionality or improvement` - Blue
- Actions
  - General
    - Runners > `All repositories`
    - Artifact and log retention > `30 days`
    - Approval for running fork pull request workflows from contributors > `Require approval for all external contributors`
  - Fork pull request workflows in private repositories
    - Run workflows from fork pull requests > `Disable`
  - Workflow permissions > `Read repository contents and packages permissions`
  - Allow GitHub Actions to create and approve pull requests > `Enable`
- Discussions > Set ad-hoc repository
- Packages
  - Package creation > `Private only`
  - Default package settings > `Inherit access from source repository`
- Authentication security
  - Require two-factor authentication for everyone in the kemadev organization > `Enable`
    - Only allow secure two-factor methods > `Enable`
- Deploy keys > `Disable`
- Code security
  - Configurations
    - New configuration
      - Everything > `Enable` / `Enforce` / `All repositories`
      - Code scanning > `Disable` (we use another tool)
      - Secret scanning
        - Scan for generic passwords > `Disable` (not supported)
        - Push protection
          - Bypass privileges
            - Select actors
              - `Maintain role`
              - `Repository admin role`
      - Code scanning
        - Runner type > `Standard`
  - Global settings
    - GitHub presets
      - Dismiss low-impact alerts for development-scoped dependencies > `Disable`
- Verified and approved domains > Verify domains
- OAuth app policy > `Access restricted`
- Personal access tokens
  - Enroll
  - Fine-grained personal access tokens > `Allow access via fine-grained personal access tokens`
  - Require approval of fine-grained personal access tokens > `Require administrator approval`
  - Set maximum lifetimes for personal access tokens
    - Fine-grained personal access tokens must expire > `90 days`
