package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

type Projects struct {
	ID           int
	Title        string
	Author       string
	StartDate    string
	EndDate      string
	Duration     string
	DescProjects string
	NodeJS       bool
	ReactJS      bool
	NextJS       bool
	TypeScript   bool
	Image        string
}

var dataProjects = []Projects{
	{
		Title:        "Ameno ameno latire Latiremo Dori me",
		Author:       "Nafiisan N. Achmad",
		StartDate:    "2023-06-06",
		EndDate:      "2023-06-07",
		Duration:     "durasi : 4 bulan",
		DescProjects: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin quis risus ut mi euismod sodales. Mauris id quam ut massa sodales faucibus consectetur sit amet dolor. ",
		NodeJS:       true,
		ReactJS:      true,
		NextJS:       true,
		TypeScript:   true,
	},

	{
		Title:        "Ameno ameno latire Latiremo Dori me",
		Author:       "Nero002",
		StartDate:    "2023-06-06",
		EndDate:      "2023-06-07",
		Duration:     "durasi : 4 bulan",
		DescProjects: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin quis risus ut mi euismod sodales. Mauris id quam ut massa sodales faucibus consectetur sit amet dolor. ",
		NodeJS:       true,
		ReactJS:      true,
		NextJS:       true,
		TypeScript:   true,
	},
}

func main() {

	e := echo.New()
	e.Static("/assets", "assets")

	e.GET("/", Home)
	e.GET("/contactMe", contactMe)
	e.GET("/project", createProject)
	e.GET("/projectDetail/:id", projectDetail)

	e.POST("/add-project", addProject)
	// e.POST("/edit-project/:id", editProject)
	e.POST("/delete-project/:id", deleteProject)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

// home
func Home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	post := map[string]interface{}{
		"Project": dataProjects,
	}

	return tmpl.Execute(c.Response(), post)
}

// contact
func contactMe(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact-me.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return tmpl.Execute(c.Response(), nil)
}

// project page
func createProject(c echo.Context) error {

	var tmpl, err = template.ParseFiles("views/project.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

// detail time
func countingDuration(startDate, endDate string) string {
	startTime, _ := time.Parse("2006-01-02", startDate)
	endTime, _ := time.Parse("2006-01-02", endDate)

	durationTime := int(endTime.Sub(startTime).Hours())
	durationDays := durationTime / 24
	durationWeeks := durationDays / 7
	durationMonths := durationWeeks / 4
	durationYears := durationMonths / 12
	var duration string
	if durationYears > 1 {
		duration = strconv.Itoa(durationYears) + " years"
	} else if durationYears > 0 {
		duration = strconv.Itoa(durationYears) + " year"
	} else {
		if durationMonths > 1 {
			duration = strconv.Itoa(durationMonths) + " months"
		} else if durationMonths > 0 {
			duration = strconv.Itoa(durationMonths) + " month"
		} else {
			if durationWeeks > 1 {
				duration = strconv.Itoa(durationWeeks) + " weeks"
			} else if durationWeeks > 0 {
				duration = strconv.Itoa(durationWeeks) + " week"
			} else {
				if durationDays > 1 {
					duration = strconv.Itoa(durationDays) + " days"
				} else {
					duration = strconv.Itoa(durationDays) + " day"
				}
			}
		}
	}
	return duration
}

// add project
func addProject(c echo.Context) error {
	title := c.FormValue("inputTitle")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	duration := countingDuration(startDate, endDate)
	DescProjects := c.FormValue("inputDescription")

	var nodeJS bool
	if c.FormValue("nodeJS") == "yes" {
		nodeJS = true
	}
	var nextJS bool
	if c.FormValue("nextJS") == "yes" {
		nextJS = true
	}
	var reactJS bool
	if c.FormValue("reactJS") == "yes" {
		reactJS = true
	}
	var typeScript bool
	if c.FormValue("typeScript") == "yes" {
		typeScript = true
	}

	image := c.FormValue("inputImage")

	var addProject = Projects{
		Title:        title,
		Author:       "Anonymous",
		StartDate:    startDate,
		EndDate:      endDate,
		Duration:     duration,
		DescProjects: DescProjects,
		NodeJS:       nodeJS,
		ReactJS:      nextJS,
		NextJS:       reactJS,
		TypeScript:   typeScript,
		Image:        image,
	}

	dataProjects = append(dataProjects, addProject)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

// project detail
func projectDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var tmpl, err = template.ParseFiles("views/project-detail.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	var listProjects = Projects{}
	for index, data := range dataProjects {
		if id == index {
			listProjects = Projects{
				Title:        data.Title,
				Author:       data.Author,
				StartDate:    data.StartDate,
				EndDate:      data.EndDate,
				Duration:     data.Duration,
				DescProjects: data.DescProjects,
				NodeJS:       data.NodeJS,
				ReactJS:      data.ReactJS,
				NextJS:       data.NextJS,
				TypeScript:   data.TypeScript,
			}
		}
	}

	data := map[string]interface{}{
		"Project": listProjects,
	}

	return tmpl.Execute(c.Response(), data)
}

// edit project
// func editProject() error {
// 	id, _ := strconv.Atoi(c.Param("id"))

// 	return c.Redirect(http.StatusMovedPermanently, "/")
// }

// delete project
func deleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println("Index : ", id)

	dataProjects = append(dataProjects[:id], dataProjects[id+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/")
}
