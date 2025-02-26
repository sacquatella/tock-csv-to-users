// Copyright © 2025 Acquatella Stephan
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	var inputFile string
	var outputFile string

	var rootCmd = &cobra.Command{
		Use:   "tock-csv-to-users",
		Short: "Build de Tock auth configmap from a CSV file",
		Run: func(cmd *cobra.Command, args []string) {
			convertCSVToYAML(inputFile, outputFile)
		},
	}

	// Add flags -f et -o
	rootCmd.Flags().StringVarP(&inputFile, "file", "f", "", "pathto csv file(mandatory)")
	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "output.yaml", "output configmap file")
	rootCmd.MarkFlagRequired("file")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Erreur :", err)
		os.Exit(1)
	}
}

func convertCSVToYAML(inputFile, outputFile string) {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Can't open csv file:", err)
		return
	}
	defer file.Close()

	// Read csv file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Can't read csv file :", err)
		return
	}

	// Init fields for ConfigMap
	var emails, passwords, organizations, roles []string

	for i, record := range records {
		if i == 0 {
			continue
		}
		emails = append(emails, record[0])
		passwords = append(passwords, record[1])
		organizations = append(organizations, record[2])
		roles = append(roles, record[3])
	}

	// Build ConfigMap YAML
	configMap := fmt.Sprintf(`apiVersion: v1
kind: ConfigMap
metadata:
  name: admin-web-auth-cfg
  labels:
    app.kubernetes.io/name: admin-web-auth
    app.kubernetes.io/component: admin-web
data:
  tock_users: "%s"
  tock_passwords: "%s"
  tock_organizations: "%s"
  tock_roles: "%s"
`, strings.Join(emails, ","), strings.Join(passwords, ","), strings.Join(organizations, ","), strings.Join(roles, ","))

	// Write ConfigMap
	err = os.WriteFile(outputFile, []byte(configMap), 0644)
	if err != nil {
		fmt.Println("Cant' write YAML file:", err)
		return
	}

	fmt.Printf("YAML file build with success : %s\n", outputFile)
}
