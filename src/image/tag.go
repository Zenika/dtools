// dtools
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: src/image/tag.go
// Original timestamp: 2023/11/27 20:36

package image

import (
	"context"
	"dtools/auth"
	"fmt"
	cerr "github.com/jeanfrancoisgratton/customError"
	hf "github.com/jeanfrancoisgratton/helperFunctions"
)

func Tag(sourceTag, newTag string) *cerr.CustomError {
	tExsts := false
	var ce *cerr.CustomError
	cli := auth.ClientConnect(true)

	if tExsts, ce = TagExists(cli, newTag); ce != nil {
		return ce
	}

	// Now that we've settled the issue of error, let's concentrate on the outcome
	if !OverwriteTag && tExsts {
		return &cerr.CustomError{Title: "Could not write tag",
			Message: fmt.Sprintf("Tag %s exists and 'overwritetag' is set to false",
				hf.Blue(newTag))}
	}

	// ... and now we tag
	err := cli.ImageTag(context.Background(), sourceTag, newTag)
	if err != nil {
		return &cerr.CustomError{Title: "Error tagging image: ", Message: err.Error()}
	}
	fmt.Printf("%s %s to %s\n", hf.Green("Successfully tagged"), hf.White(sourceTag),
		hf.White(newTag))
	return nil
}
