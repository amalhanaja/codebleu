# Codebleu: Powerful Pull Request/Merge Request Reviewer 

Streamline your code review workflow with Codebleu, a powerful tool that assists you in reviewing pull request (PR) or merge request (MR) changes. Leveraging Gemini AI, Codebleu provides valuable insights and potential issues to consider, empowering you to conduct more efficient and effective code reviews.

```bash
NAME:
   Codebleu - Review PR / MR Diff Changes

USAGE:
   Codebleu [global options] command [command options] 

VERSION:
   v0.0.1

DESCRIPTION:
   Pull Request / Merge Request reviewer agent

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --model value, -m value       uses model to review pull request (options: "gemini-1.5-flash" (default), "gemini-1.5-pro", "gemini-1.0-pro") (default: "gemini-1.5-flash") [$MODEL]
   --repository value, -r value  hosted remote repository provider name (options: "bitbucket", "github") [$REPOSITORY_PROVIDER]
   --id value                    pull request id [$PULL_REQUEST_ID]
   --system-instruction value    Custom system instruction for review pull request diff chages [$SYSTEM_INSTRUCTION]
   --help, -h                    show help
   --version, -v                 print the version
```

### Example
- **Bitbucket**: [Pull Request](https://bitbucket.org/amalhanaja/test/pull-requests/8)

- **Github**: [Pull Request](https://github.com/amalhanaja/codebleu/pull/1)

## [Get Codebleu](https://github.com/amalhanaja/codebleu/releases)

## Usage

### Github

#### Configure Environment Variable

```bash
export GH_ACCESS_TOKEN=<YOUR_GITHUB_ACCESS_TOKEN>
export GH_OWNER=<YOUR_GITHUB_REPOSITORY_OWNER>
export GH_REPO_SLUG=<YOUR_REPOSITORY_SLG>
export GEMINI_API_KEY=<YOUR_GEMINI_API_KEY>
```
#### Command

```bash
codebleu --repository github --id <PULL_REQUEST_ID>
```

#### Example:

Pull Request Url to review : https://github.com/amalhanaja/codebleu/pull/1

**Environment Variable**
```bash
export GH_ACCESS_TOKEN="secret_access_token"
export GH_OWNER="amalhanaja"
export GH_REPO_SLUG="codebleu"
export GEMINI_API_KEY="secret_api_key"
```

**Command**
```bash
codebleu --repository github --id 1
```

### Bitbucket

#### Configure Environment Variable

```bash
export BITBUCKET_ACCESS_TOKEN=<YOUR_BITBUCKET_ACCESS_TOKEN>
export BITBUCKET_REPO_SLUG=<YOUR_REPOSITORY_SLUG>
export BITBUCKET_WORKSPACE=<YOUR_REPOSITORY_WORKSPACE>
export GEMINI_API_KEY=<YOUR_GEMINI_API_KEY>
```
#### Command

```bash
codebleu --repository github --id <PULL_REQUEST_ID>
```

#### Example:

Pull Request Url to review : https://bitbucket.org/amalhanaja/test/pull-requests/8

**Environment Variable**
```bash
export BITBUCKET_ACCESS_TOKEN="secret_access_token"
export BITBUCKET_WORKSPACE="amalhanaja"
export BITBUCKET_REPO_SLUG="test"
export GEMINI_API_KEY="secret_api_key"
```

**Command**
```bash
codebleu --repository bitbucket --id 8
```