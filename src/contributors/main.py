from datetime import datetime, timedelta, timezone

import json
from math import floor
from pydriller import Repository
import sys

if len(sys.argv) > 1:
    for i, arg in enumerate(sys.argv):
        if i == 0:
            continue
else:
    print("No arguments were passed other than the script name.")

repositoryAlias  = sys.argv[1]
repositoryName = sys.argv[2]

repositoryUrl = "../../repos/" + repositoryAlias

outputJson = {
  'repositoryAlias': repositoryAlias,
  'repositoryName': repositoryName
}

contributors = set([])

# dt2 = datetime.now()
# dt1 = dt2 - timedelta(days=365 * 100)

for commit in Repository(repositoryUrl).traverse_commits():
  contributors.add(commit.author.name)

outputJson["contributors"] = len(contributors)

print(json.dumps(outputJson))