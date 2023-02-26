// SPDX-FileCopyrightText: 2023 Dinesh Ravi
//
// SPDX-License-Identifier: Apache-2.0

package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/alexeyco/simpletable"
)

type SPDX struct {
	Spdxid                     string                       `json:"SPDXID,omitempty"`
	SpdxVersion                string                       `json:"spdxVersion,omitempty"`
	CreationInfo               CreationInfo                 `json:"creationInfo,omitempty"`
	Name                       string                       `json:"name,omitempty"`
	DataLicense                string                       `json:"dataLicense,omitempty"`
	HasExtractedLicensingInfos []HasExtractedLicensingInfos `json:"hasExtractedLicensingInfos,omitempty"`
	DocumentNamespace          string                       `json:"documentNamespace,omitempty"`
	DocumentDescribes          []string                     `json:"documentDescribes,omitempty"`
	Packages                   []Packages                   `json:"packages,omitempty"`
	Files                      []Files                      `json:"files,omitempty"`
	Relationships              []Relationships              `json:"relationships,omitempty"`
}
type CreationInfo struct {
	Created            time.Time `json:"created,omitempty"`
	Creators           []string  `json:"creators,omitempty"`
	LicenseListVersion string    `json:"licenseListVersion,omitempty"`
}
type HasExtractedLicensingInfos struct {
	LicenseID     string `json:"licenseId,omitempty"`
	ExtractedText string `json:"extractedText,omitempty"`
	Name          string `json:"name,omitempty"`
}

type ExternalRefs struct {
	ReferenceCategory string `json:"referenceCategory,omitempty"`
	ReferenceLocator  string `json:"referenceLocator,omitempty"`
	ReferenceType     string `json:"referenceType,omitempty"`
}
type Packages struct {
	Spdxid           string         `json:"SPDXID,omitempty"`
	CopyrightText    string         `json:"copyrightText,omitempty"`
	DownloadLocation string         `json:"downloadLocation,omitempty"`
	ExternalRefs     []ExternalRefs `json:"externalRefs,omitempty"`
	FilesAnalyzed    bool           `json:"filesAnalyzed,omitempty"`
	Homepage         string         `json:"homepage,omitempty"`
	LicenseConcluded string         `json:"licenseConcluded,omitempty"`
	LicenseDeclared  string         `json:"licenseDeclared,omitempty"`
	Name             string         `json:"name,omitempty"`
	Supplier         string         `json:"supplier,omitempty"`
	VersionInfo      string         `json:"versionInfo,omitempty"`
	HasFiles         []string       `json:"hasFiles,omitempty"`
}
type Checksums struct {
	Algorithm     string `json:"algorithm,omitempty"`
	ChecksumValue string `json:"checksumValue,omitempty"`
}
type Files struct {
	Spdxid             string      `json:"SPDXID,omitempty"`
	Checksums          []Checksums `json:"checksums,omitempty"`
	CopyrightText      string      `json:"copyrightText,omitempty"`
	FileName           string      `json:"fileName,omitempty"`
	LicenseConcluded   string      `json:"licenseConcluded,omitempty"`
	LicenseInfoInFiles []string    `json:"licenseInfoInFiles,omitempty"`
}
type Relationships struct {
	SpdxElementID      string `json:"spdxElementId,omitempty"`
	RelatedSpdxElement string `json:"relatedSpdxElement,omitempty"`
	RelationshipType   string `json:"relationshipType,omitempty"`
}

func (t *SPDX) Load(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}
	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}

	return nil
}

func (s *SPDX) PrintMeta() {

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Key"},
			{Align: simpletable.AlignCenter, Text: "Value"},
		},
	}

	var cells [][]*simpletable.Cell

	idx := 0

	idx++

	cells = append(cells, []*simpletable.Cell{

		{Text: fmt.Sprintf("%d", idx)},
		{Text: blue("Spdx ID")},
		{Text: s.Spdxid},
	})
	idx++
	cells = append(cells, []*simpletable.Cell{
		{Text: fmt.Sprintf("%d", idx)},
		{Text: blue("Spdx version")},
		{Text: s.SpdxVersion},
	})
	idx++
	cells = append(cells, []*simpletable.Cell{
		{Text: fmt.Sprintf("%d", idx)},
		{Text: blue("Spdx creation date")},
		{Text: yellow(s.CreationInfo.Created.Format(time.RFC822))},
	})
	idx++
	if len(s.CreationInfo.Creators) > 0 {
		cells = append(cells, []*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: blue("created by")},
			{Text: strings.Join(s.CreationInfo.Creators, ", ")},
		})
	}
	idx++
	cells = append(cells, []*simpletable.Cell{
		{Text: fmt.Sprintf("%d", idx)},
		{Text: blue("Project Name")},
		{Text: red(s.Name)},
	})
	idx++
	cells = append(cells, []*simpletable.Cell{
		{Text: fmt.Sprintf("%d", idx)},
		{Text: blue("File License(not projects)")},
		{Text: s.DataLicense},
	})
	idx++
	cells = append(cells, []*simpletable.Cell{
		{Text: fmt.Sprintf("%d", idx)},
		{Text: blue("Document Namespace")},
		{Text: s.DocumentNamespace},
	})
	idx++
	cells = append(cells, []*simpletable.Cell{
		{Text: fmt.Sprintf("%d", idx)},
		{Text: blue("Document Describes")},
		{Text: strings.Join(s.DocumentDescribes, ", ")},
	})
	idx++
	cells = append(cells, []*simpletable.Cell{
		{Text: fmt.Sprintf("%d", idx)},
		{Text: blue("Number of Packages")},
		{Text: red(fmt.Sprintf("%d", len(s.Packages)))},
	})
	idx++
	cells = append(cells, []*simpletable.Cell{
		{Text: fmt.Sprintf("%d", idx)},
		{Text: blue("Number of Files")},
		{Text: red(fmt.Sprintf("%d", len(s.Files)))},
	})
	idx++
	cells = append(cells, []*simpletable.Cell{
		{Text: fmt.Sprintf("%d", idx)},
		{Text: blue("Number of Relationships")},
		{Text: red(fmt.Sprintf("%d", len(s.Relationships)))},
	})

	table.Body = &simpletable.Body{Cells: cells}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}

func (s *SPDX) PrintFiles(nf int) {

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "FileName"},
			// {Align: simpletable.AlignCenter, Text: "LicenseConcluded"},
			// {Align: simpletable.AlignCenter, Text: "LicenseInfoInFiles"},
			// {Align: simpletable.AlignCenter, Text: "SPDXId"},
			// {Align: simpletable.AlignCenter, Text: "CopyrightText"},
			{Align: simpletable.AlignCenter, Text: "checksum"},
			{Align: simpletable.AlignCenter, Text: "Algorithm"},
		},
	}
	// {Align: simpletable.AlignCenter, Text: "Checksums"},
	// {Align: simpletable.AlignCenter, Text: "Algorithm - Checksums"},
	files := s.Files
	var file Files
	var n int
	// var licenseinfo string
	lenFiles := len(files)
	// fmt.Println(lenFiles)
	var cells [][]*simpletable.Cell

	for id := 0; id < nf; id++ {

		file = files[id]

		n = id + 1
		// licenseinfo = strings.Join(file.LicenseInfoInFiles, ", ")

		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", n)},
			{Text: file.FileName},
			// {Text: file.LicenseConcluded},
			// {Text: licenseinfo},
			// {Text: file.Spdxid},
			// {Text: file.CopyrightText},
			{Text: file.Checksums[0].ChecksumValue},
			{Text: file.Checksums[0].Algorithm},
		})

	}
	table.Body = &simpletable.Body{Cells: cells}

	if lenFiles > 0 {
		table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Span: 4, Text: blue(fmt.Sprintf("There are %d Files", lenFiles))},
		}}
	}
	table.SetStyle(simpletable.StyleUnicode)
	table.Println()
}
func (s *SPDX) PrintFilesIP(nf int) {

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "FileName"},
			{Align: simpletable.AlignCenter, Text: "LicenseConcluded"},
			{Align: simpletable.AlignCenter, Text: "LicenseInfoInFiles"},
			// {Align: simpletable.AlignCenter, Text: "SPDXId"},
			{Align: simpletable.AlignCenter, Text: "CopyrightText"},
			// {Align: simpletable.AlignCenter, Text: "checksum"},
			// {Align: simpletable.AlignCenter, Text: "Algorithm"},
		},
	}
	// {Align: simpletable.AlignCenter, Text: "Checksums"},
	// {Align: simpletable.AlignCenter, Text: "Algorithm - Checksums"},
	files := s.Files
	var file Files
	var n int
	var licenseinfo string
	lenFiles := len(files)
	// fmt.Println(lenFiles)
	var cells [][]*simpletable.Cell

	for id := 0; id < nf; id++ {

		file = files[id]

		n = id + 1
		licenseinfo = strings.Join(file.LicenseInfoInFiles, ", ")

		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", n)},
			{Text: file.FileName},
			{Text: file.LicenseConcluded},
			{Text: licenseinfo},
			// {Text: file.Spdxid},
			{Text: file.CopyrightText},
			// {Text: file.Checksums[0].ChecksumValue},
			// {Text: file.Checksums[0].Algorithm},
		})

	}
	table.Body = &simpletable.Body{Cells: cells}

	if lenFiles > 0 {
		table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Span: 5, Text: blue(fmt.Sprintf("There are %d Files", lenFiles))},
		}}
	}
	table.SetStyle(simpletable.StyleUnicode)
	table.Println()
}
func (s *SPDX) PrintFilesExt(nf int) {

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "FileName"},
			{Align: simpletable.AlignCenter, Text: "LicenseConcluded"},
			{Align: simpletable.AlignCenter, Text: "LicenseInfoInFiles"},
			{Align: simpletable.AlignCenter, Text: "SPDXId"},
			{Align: simpletable.AlignCenter, Text: "CopyrightText"},
			{Align: simpletable.AlignCenter, Text: "checksum"},
			{Align: simpletable.AlignCenter, Text: "Algorithm"},
		},
	}
	// {Align: simpletable.AlignCenter, Text: "Checksums"},
	// {Align: simpletable.AlignCenter, Text: "Algorithm - Checksums"},
	files := s.Files
	var file Files
	var n int
	var licenseinfo string
	lenFiles := len(files)
	// fmt.Println(lenFiles)
	var cells [][]*simpletable.Cell

	for id := 0; id < nf; id++ {

		file = files[id]

		n = id + 1
		licenseinfo = strings.Join(file.LicenseInfoInFiles, ", ")

		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", n)},
			{Text: file.FileName},
			{Text: file.LicenseConcluded},
			{Text: licenseinfo},
			{Text: file.Spdxid},
			{Text: file.CopyrightText},
			{Text: file.Checksums[0].ChecksumValue},
			{Text: file.Checksums[0].Algorithm},
		})

	}
	table.Body = &simpletable.Body{Cells: cells}

	if lenFiles > 0 {
		table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Span: 8, Text: blue(fmt.Sprintf("There are %d Files", lenFiles))},
		}}
	}
	table.SetStyle(simpletable.StyleUnicode)
	table.Println()
}
func (s *SPDX) Printpkgs(np int) {

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Supplier"},
			{Align: simpletable.AlignCenter, Text: "Name"},
			{Align: simpletable.AlignCenter, Text: "VersionInfo"},
			{Align: simpletable.AlignCenter, Text: "Homepage"},
			// {Align: simpletable.AlignCenter, Text: "LicenseDeclared"},
			// {Align: simpletable.AlignCenter, Text: "LicenseConcluded"},
			// {Align: simpletable.AlignCenter, Text: "FilesAnalyzed"},
			// {Align: simpletable.AlignCenter, Text: "DownloadLocation"},
			// {Align: simpletable.AlignCenter, Text: "CopyrightText"},
			// {Align: simpletable.AlignCenter, Text: "Spdxid"},
		},
	}
	// {Align: simpletable.AlignCenter, Text: "Checksums"},
	// {Align: simpletable.AlignCenter, Text: "Algorithm - Checksums"},
	pkgs := s.Packages
	var pkg Packages
	var n int
	// var licenseinfo string
	lenPkgs := len(pkgs)
	// fmt.Println(lenPkgs)
	var cells [][]*simpletable.Cell

	for id := 0; id < np; id++ {

		pkg = pkgs[id]

		n = id + 1
		// licenseinfo = strings.Join(file.LicenseInfoInFiles, ", ")

		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", n)},
			{Text: pkg.Supplier},
			{Text: pkg.Name},
			{Text: pkg.VersionInfo},
			{Text: pkg.Homepage},
			// {Text: pkg.LicenseDeclared},
			// {Text: pkg.LicenseConcluded},
			// {Text: fmt.Sprintf("%v", pkg.FilesAnalyzed)},
			// {Text: pkg.DownloadLocation},
			// {Text: pkg.CopyrightText},
			// {Text: pkg.Spdxid},
		})

	}
	table.Body = &simpletable.Body{Cells: cells}

	if lenPkgs > 0 {
		table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Span: 5, Text: blue(fmt.Sprintf("There are %d pkgs", lenPkgs))},
		}}
	}
	table.SetStyle(simpletable.StyleUnicode)
	table.Println()

}

func (s *SPDX) PrintpkgsIP(np int) {

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			// {Align: simpletable.AlignCenter, Text: "Supplier"},
			{Align: simpletable.AlignCenter, Text: "Name"},
			{Align: simpletable.AlignCenter, Text: "VersionInfo"},
			// {Align: simpletable.AlignCenter, Text: "Homepage"},
			{Align: simpletable.AlignCenter, Text: "LicenseDeclared"},
			{Align: simpletable.AlignCenter, Text: "LicenseConcluded"},
			{Align: simpletable.AlignCenter, Text: "FilesAnalyzed"},
			// {Align: simpletable.AlignCenter, Text: "DownloadLocation"},
			{Align: simpletable.AlignCenter, Text: "CopyrightText"},
			{Align: simpletable.AlignCenter, Text: "FilesAnalyzed"},
			// {Align: simpletable.AlignCenter, Text: "Spdxid"},
		},
	}
	// {Align: simpletable.AlignCenter, Text: "Checksums"},
	// {Align: simpletable.AlignCenter, Text: "Algorithm - Checksums"},
	pkgs := s.Packages
	var pkg Packages
	var n int
	// var licenseinfo string
	lenPkgs := len(pkgs)
	// fmt.Println(lenPkgs)
	var cells [][]*simpletable.Cell

	for id := 0; id < np; id++ {

		pkg = pkgs[id]

		n = id + 1
		// licenseinfo = strings.Join(file.LicenseInfoInFiles, ", ")

		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", n)},
			// {Text: pkg.Supplier},
			{Text: pkg.Name},
			{Text: pkg.VersionInfo},
			// {Text: pkg.Homepage},
			{Text: pkg.LicenseDeclared},
			{Text: pkg.LicenseConcluded},

			// {Text: pkg.DownloadLocation},
			{Text: pkg.CopyrightText},
			{Text: fmt.Sprintf("%v", pkg.FilesAnalyzed)},
			// {Text: pkg.Spdxid},
		})

	}
	table.Body = &simpletable.Body{Cells: cells}

	if lenPkgs > 0 {
		table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Span: 8, Text: blue(fmt.Sprintf("There are %d pkgs", lenPkgs))},
		}}
	}
	table.SetStyle(simpletable.StyleUnicode)
	table.Println()

}
func (s *SPDX) PrintpkgsExt(np int) {

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Supplier"},
			{Align: simpletable.AlignCenter, Text: "Name"},
			{Align: simpletable.AlignCenter, Text: "VersionInfo"},
			{Align: simpletable.AlignCenter, Text: "Homepage"},
			{Align: simpletable.AlignCenter, Text: "LicenseDeclared"},
			{Align: simpletable.AlignCenter, Text: "LicenseConcluded"},
			{Align: simpletable.AlignCenter, Text: "FilesAnalyzed"},
			{Align: simpletable.AlignCenter, Text: "DownloadLocation"},
			{Align: simpletable.AlignCenter, Text: "CopyrightText"},
			// {Align: simpletable.AlignCenter, Text: "Spdxid"},
		},
	}
	// {Align: simpletable.AlignCenter, Text: "Checksums"},
	// {Align: simpletable.AlignCenter, Text: "Algorithm - Checksums"},
	pkgs := s.Packages
	var pkg Packages
	var n int
	// var licenseinfo string
	lenPkgs := len(pkgs)
	// fmt.Println(lenPkgs)
	var cells [][]*simpletable.Cell

	for id := 0; id < np; id++ {

		pkg = pkgs[id]

		n = id + 1
		// licenseinfo = strings.Join(file.LicenseInfoInFiles, ", ")

		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", n)},
			{Text: pkg.Supplier},
			{Text: pkg.Name},
			{Text: pkg.VersionInfo},
			{Text: pkg.Homepage},
			{Text: pkg.LicenseDeclared},
			{Text: pkg.LicenseConcluded},
			{Text: fmt.Sprintf("%v", pkg.FilesAnalyzed)},
			{Text: pkg.DownloadLocation},
			{Text: pkg.CopyrightText},
			// {Text: pkg.Spdxid},
		})

	}
	table.Body = &simpletable.Body{Cells: cells}

	if lenPkgs > 0 {
		table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Span: 10, Text: blue(fmt.Sprintf("There are %d pkgs", lenPkgs))},
		}}
	}
	table.SetStyle(simpletable.StyleUnicode)
	table.Println()

}

func (s *SPDX) PrintRels(np int) {

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "SpdxElementID"},
			{Align: simpletable.AlignCenter, Text: "RelationshipType"},
			{Align: simpletable.AlignCenter, Text: "RelatedSpdxElement"},
		},
	}
	// {Align: simpletable.AlignCenter, Text: "Checksums"},
	// {Align: simpletable.AlignCenter, Text: "Algorithm - Checksums"},
	rels := s.Relationships
	var rel Relationships
	var n int
	// var licenseinfo string
	lenrels := len(rels)
	// fmt.Println(lenrels)
	var cells [][]*simpletable.Cell

	for id := 0; id < np; id++ {

		rel = rels[id]

		n = id + 1
		// licenseinfo = strings.Join(file.LicenseInfoInFiles, ", ")
		switch rt := rel.RelationshipType; rt {
		case "CONTAINS":
			cells = append(cells, *&[]*simpletable.Cell{
				{Text: fmt.Sprintf("%d", n)},
				{Text: rel.SpdxElementID},
				{Text: blue(rt)},
				{Text: rel.RelatedSpdxElement},
			})
		case "DEPENDS_ON":
			cells = append(cells, *&[]*simpletable.Cell{
				{Text: fmt.Sprintf("%d", n)},
				{Text: rel.SpdxElementID},
				{Text: yellow(rt)},
				{Text: rel.RelatedSpdxElement},
			})
		case "DESCRIBES":
			cells = append(cells, *&[]*simpletable.Cell{
				{Text: fmt.Sprintf("%d", n)},
				{Text: rel.SpdxElementID},
				{Text: red(rt)},
				{Text: rel.RelatedSpdxElement},
			})
		default:
			cells = append(cells, *&[]*simpletable.Cell{
				{Text: fmt.Sprintf("%d", n)},
				{Text: rel.SpdxElementID},
				{Text: gray(rt)},
				{Text: rel.RelatedSpdxElement},
			})

		}

	}
	table.Body = &simpletable.Body{Cells: cells}

	if lenrels > 0 {
		table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Span: 4, Text: blue(fmt.Sprintf("There are %d relationships", lenrels))},
		}}
	}
	table.SetStyle(simpletable.StyleUnicode)
	table.Println()

}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

//	func (s *SPDX) getrelspdxelem(spdxID string) (string, Relationships) {
//		rels := s.Relationships
//		for _, rel := range rels {
//			if rel.SpdxElementID == spdxID {
//				return rel.RelatedSpdxElement, rel
//			}
//		}
//		return "", Relationships{}
//	}
func (s *SPDX) getdependson(d int, rel1 Relationships) {
	pkgs := s.Packages
	var nextSPDXidDetail string
	var SPDXidDetail string
	// fmt.Println("===================DEPENDS_ON======================")
	for _, pkg := range pkgs {
		if pkg.Spdxid == rel1.SpdxElementID {
			SPDXidDetail = fmt.Sprintf("%v %v version: %v", blue(pkg.Name), yellow("|"), blue(pkg.VersionInfo))
		}
		if pkg.Spdxid == rel1.RelatedSpdxElement {
			nextSPDXidDetail = fmt.Sprintf("%v %v version: %v", blue(pkg.Name), yellow("|"), blue(pkg.VersionInfo))
		}

	}

	fmt.Println(fmt.Sprintf("%v =====> %v Pkg %v =====> %v", SPDXidDetail, green("DEPENDS_ON"), blue(fmt.Sprintf("%d", d)), nextSPDXidDetail))
}
func (s *SPDX) getspdxpkg(d int, spdxID string, rel1 Relationships) {
	pkgs := s.Packages
	var filespdxids []string
	var SPDXidDetail string
	// var nextSPDXid string
	// var rel Relationships
	for _, pkg := range pkgs {
		if pkg.Spdxid == spdxID {
			SPDXidDetail = fmt.Sprintf("%v | version: %v", yellow(pkg.Name), yellow(pkg.VersionInfo))
			if len(pkg.HasFiles) > 0 {

				filespdxids = pkg.HasFiles
				for i, filespdx := range filespdxids {
					s.getspdxfile(i, filespdx, SPDXidDetail)
				}
			} else {
				var f int
				for _, rel1 := range s.Relationships {
					if rel1.SpdxElementID == spdxID {
						switch rt := rel1.RelationshipType; rt {
						case "CONTAINS":

							s.getspdxfile(f, rel1.RelatedSpdxElement, SPDXidDetail)
						}
					}
				}

			}

			// d++
			// nextSPDXid, rel = s.getrelspdxelem(spdxID)
			// if nextSPDXid != "" {
			// 	s.getspdxpkg(d, nextSPDXid, rel)
			// }
		}
	}
}
func (s *SPDX) getspdxfile(i int, spdxID string, SPDXidDetail string) {
	files := s.Files
	for _, file := range files {
		if file.Spdxid == spdxID {
			i++
			fmt.Println(green(fmt.Sprintf("%v %v %v File %v %v %v", SPDXidDetail, yellow("---->"), red("CONTAINS"), blue(fmt.Sprintf("%d", i)), yellow("---->"), red(file.FileName))))
		}

	}
	// n = 0
	// return ""
}

func (s *SPDX) PrintDigRels() {
	rels := s.Relationships
	var rel Relationships
	var n int
	_ = n
	lenrels := len(rels)
	var spdxids []string = make([]string, lenrels)
	// var licenseinfo string
	//
	for id := 0; id < lenrels; id++ {
		rel = rels[id]
		spdxids = append(spdxids, rel.SpdxElementID)
	}
	unique_spdxIDs := removeDuplicateStr(spdxids)
	len_unique_spdxIDs := len(unique_spdxIDs)
	fmt.Println("No. of unique spdxID :", len_unique_spdxIDs)
	// _ = len_unique_spdxIDs
	var d int

	fmt.Println(red("===================DESCRIBES======================"))
	for _, rel1 := range rels {
		switch rt := rel1.RelationshipType; rt {
		case "DESCRIBES":
			d++
			s.getspdxpkg(d, rel1.RelatedSpdxElement, rel1)
		}

	}
	d = 0
	fmt.Println(red("===================DEPENDS_ON======================"))
	for _, rel1 := range rels {
		switch rt := rel1.RelationshipType; rt {
		case "DEPENDS_ON":
			d++
			s.getdependson(d, rel1)
			s.getspdxpkg(d, rel1.RelatedSpdxElement, rel1)
		}
	}

}
