workflow "Conform Pull Request" {
  on       = "pull_request"

  resolves = [
    "conform"
  ]
}

action "conform" {
  uses    = "docker://autonomy/conform:v0.1.0-alpha.13"

  secrets = [
    "GITHUB_TOKEN"
  ]
}