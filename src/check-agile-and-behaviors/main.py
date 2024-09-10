from datetime import datetime, timedelta, timezone

import json
from math import floor
from pydriller import Repository
import hashlib
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

# dt2 = datetime(2024, 1, 1, 0, 0, 0, tzinfo=timezone.utc)
dt2 = datetime.now()
dt1 = dt2 - timedelta(days=365)

commits = []
contributors = set([])
contributorsInfo = {}

for commit in Repository(repositoryUrl, since=dt1, to=dt2).traverse_commits():

    if(commit.in_main_branch):
      commits.append(commit.committer_date)

      contributorHash = hashlib.sha256(commit.author.name.encode()).hexdigest()

      contributors.add(contributorHash)

      if contributorHash not in contributorsInfo:
        contributorsInfo[contributorHash] = { }

      if 'commits' not in contributorsInfo[contributorHash]:
        contributorsInfo[contributorHash]['commits'] = []

      contributorsInfo[contributorHash]['commits'].append(commit.committer_date)

outputJson["contributors"] = len(contributors)

outputJson["commits"] = len(commits)
hasCommitsInInterval = True if len(commits) > 0 else False
outputJson["hasCommitsInInterval"] = hasCommitsInInterval
outputJson["lastCommitDateInterval"] = commits[0] if hasCommitsInInterval else None


intervalos = [(commits[i] - commits[i-1]).total_seconds() for i in range(1, len(commits))]
media_intervalos = sum(intervalos) / len(intervalos) if len(intervalos) > 0 else 0
outputJson["commitsIntervalInDays"] =  floor(media_intervalos / 86400 * 100) / 100

for contributorHash in contributorsInfo:
  commits = contributorsInfo[contributorHash]['commits']

  if "contributorsInfo" not in outputJson:
    outputJson["contributorsInfo"] = { }

  if contributorHash not in outputJson["contributorsInfo"]:
    outputJson["contributorsInfo"][contributorHash] = { }

  outputJson["contributorsInfo"][contributorHash]['commits'] = floor(media_intervalos / 86400 * 100) / 100

  intervalos = [(commits[i] - commits[i-1]).total_seconds() for i in range(1, len(commits))]

  media_intervalos = sum(intervalos) / len(intervalos) if len(intervalos) > 0 else 0
  outputJson["contributorsInfo"][contributorHash]['commitsIntervalInDays'] = floor(media_intervalos / 86400 * 100) / 100

def datetime_serializer(obj):
    if isinstance(obj, datetime):
        return obj.isoformat()
    raise TypeError("Type not serializable")
print(json.dumps(outputJson, default=datetime_serializer, indent=4))