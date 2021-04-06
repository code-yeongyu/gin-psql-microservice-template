from os import system
import os
from sys import argv as sys_argv
import argparse

parser = argparse.ArgumentParser(
    description="Rename the module name of gin-microservice-template.")
parser.add_argument('new_name',
                    metavar='NAME',
                    type=str,
                    help='New name for the gin-microservice-template')
args = parser.parse_args()

old_name = "gin_psql_microservice_template"
new_name = sys_argv[1]

system(
    f"grep -rli '{old_name}' * | xargs -I@ sed -i '' 's/{old_name}/{new_name}/g' @"
)

print(f"Renamed the go mod module name from {old_name} to {new_name}")