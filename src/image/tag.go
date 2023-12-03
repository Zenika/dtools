// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/image/tag.go
// Original timestamp: 2023/11/27 20:36

package image

import (
	"context"
	"dtools/auth"
	"dtools/helpers"
	"fmt"
)

func Tag(sourceTag, newTag string) error {
	cli := auth.ClientConnect(true)

	tExsts, err := TagExists(cli, newTag)
	if err != nil {
		return helpers.CustomError{Message: "Unable to tag image: " + err.Error()}
	}

	// Now that we've settled the issue of error, let's concentrate on the outcome
	if !OverwriteTag && tExsts {
		return helpers.CustomError{Message: fmt.Sprintf("Tag %s exists and 'overwritetag' is set to false",
			helpers.Blue(newTag))}
	}

	// ... and now we tag
	err = cli.ImageTag(context.Background(), sourceTag, newTag)
	if err != nil {
		return helpers.CustomError{Message: "Error tagging image: " + err.Error()}
	}
	fmt.Printf("%s %s to %s\n", helpers.Green("Successfully tagged"), helpers.White(sourceTag),
		helpers.White(newTag))
	return nil
}
