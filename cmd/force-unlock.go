/*
 * MIT License
 *
 * Copyright (c) 2017 Aurelien PERRIER
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the 'Software'), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

/*
 * Revision History:
 *     Initial: 2017/12/07      Aurelien PERRIER
 */

package cmd

import (
	"fmt"

	"github.com/mitchellh/cli"
	"github.com/perriea/tfwrapper/pkg/wrapper"
)

// FmtCommand : Struct UI Parameter
type ForceUnlockCommand struct {
	UI cli.Ui
}

// Run : Executing Terraform command
func (c *ForceUnlockCommand) Run(s []string) int {
	for _, item := range s {
		args = append(args, item)
	}

	wrapper.Action("force-unlock", args)
	c.UI.Output(fmt.Sprintf("\nIt's OK !"))
	return 0
}

// Help : Show Help
func (c *ForceUnlockCommand) Help() string {
	return "Rewrites config files to canonical format"
}

// Synopsis : Show Synopsis
func (c *ForceUnlockCommand) Synopsis() string {
	return "Rewrites config files to canonical format"
}