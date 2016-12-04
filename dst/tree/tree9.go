package autogen

/* structure
Tree
	Left
		Left
			Left
				Right
				Right
					Left
					Right
	Right
		Left
			Left
				Right
					Right
*/
// Tree : auto generated JSON container
type Tree struct {
	Left  Tree `json:"left"`
	Right Tree `json:"right"`
	Total int  `json:"total" example:"100"`
}
