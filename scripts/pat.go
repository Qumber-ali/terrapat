package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"os"
)

func main() {

	ctx := context.Background()

	var cred_options azidentity.AzureCLICredentialOptions

	reader := bufio.NewReader(os.Stdin)

	tfinput, _ := reader.ReadString('\n')

	var tfobj map[string]any

	json.Unmarshal([]byte(tfinput), &tfobj)

	tenant := tfobj["tenant"].(string)

	cred_options.TenantID = tenant

	var cropt *azidentity.AzureCLICredentialOptions

	cropt = &cred_options

	cred, err := azidentity.NewAzureCLICredential(cropt)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	var token_options policy.TokenRequestOptions

	token_options.Scopes = []string{"https://jcidata.visualstudio.com"}

	token, err := cred.GetToken(ctx, token_options)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	data := map[string]interface{}{
		"output": token.Token,
	}

	result, err := json.Marshal(data)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", result)
}
