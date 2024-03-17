package main

import (
	"fmt"
)

func main() {

	// // Historico de commits e pull requets mais recentos em repositorio
	// jsonMapInstance := map[string]string{
	// 	"query": `
	//   {
	//     repository(owner: "luiz04nl", name: "devops-ic-collector") {
	//       defaultBranchRef {
	//         target {
	//           ... on Commit {
	//             history(first: 10) {
	//               edges {
	//                 node {
	//                   committedDate
	//                   oid
	//                   messageHeadline
	//                   author {
	//                     name
	//                     date
	//                   }
	//                 }
	//               }
	//             }
	//           }
	//         }
	//       }
	//       pullRequests(last: 10, orderBy: {field: CREATED_AT, direction: DESC}) {
	//         edges {
	//           node {
	//             title
	//             state
	//             author {
	//               login
	//             }
	//             createdAt
	//           }
	//         }
	//       }
	//     }
	//   }
	// `,
	// }

	// Historico de commits e pull requets mais recentos em repositorio
	query := `
{
  search(query: "is:public stars:>=100", type: REPOSITORY, first: 10) {
    repositoryCount
    pageInfo {
      endCursor
      startCursor
    }
    edges {
      node {
        ... on Repository {
          name
          url
          stargazers {
            totalCount
          }
          collaborators(first: 3) {
            totalCount
            edges {
              permission
              node {
                id
                name
                email
              }
            }
            pageInfo {
              hasNextPage
              endCursor
            }
          }
        }
      }
    }
  }
}
`
	response, err := executeGraphQLQuery(query)
	if err != nil {
		fmt.Println(err)
		return
	}

	// repositories := gitHubSearchResponseToRepositories(response)

	fmt.Println(string(response))
}
