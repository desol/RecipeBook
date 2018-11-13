param (
  [Parameter(Mandatory=$true)][string]$server,
  [Parameter(Mandatory=$true)][string]$username,
  [Parameter(Mandatory=$true)][string] $password
)

dotnet ef dbcontext scaffold "Host=$server;Database=lhj;Username=$username;Password=$password" Npgsql.EntityFrameworkCore.PostgreSQL --output-dir Data\Models
