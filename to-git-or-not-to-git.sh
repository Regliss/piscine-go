curl -s https://01.alem.school/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{githubLogin:{_eq:\"ars192\"}}){id}}"}' | cut -d: -f4 | cut -d} -f1
