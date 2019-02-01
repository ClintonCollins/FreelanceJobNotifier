package standard_out

import (
	"FreelanceJobNotifier/models"
	"fmt"
	"strings"
)

func HandleJobGroups(group models.JobGroup) {
	fmt.Printf("###### Found %d new %s jobs ######\n", len(group.Jobs), strings.Title(group.Name))
	for _, job := range group.Jobs {
		fmt.Println(job.Title)
		fmt.Println(job.URL)
		fmt.Println()
	}
	fmt.Println()
	fmt.Printf("###### End of %s jobs ######\n", strings.Title(group.Name))
	fmt.Println()
}