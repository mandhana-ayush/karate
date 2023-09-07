// // // // // package main

// // // // // import (
// // // // // 	"assg/pizzashop/pkg/common/db"
// // // // // 	"assg/pizzashop/pkg/pizzashop"
// // // // // 	"encoding/json"
// // // // // 	"fmt"
// // // // // 	"io/ioutil"
// // // // // 	"os"

// // // // // 	"assg/pizzashop/pkg/common/models"

// // // // // 	"github.com/gin-gonic/gin"
// // // // // 	"github.com/spf13/viper"
// // // // // 	"gorm.io/gorm"
// // // // // )

// // // // // func main() {
// // // // // 	viper.SetConfigFile("../pkg/common/envs/.env")
// // // // // 	viper.ReadInConfig()

// // // // // 	port := viper.Get("PORT").(string)
// // // // // 	dbUrl := viper.Get("DB_URL").(string)

// // // // // 	fmt.Println(port, dbUrl)

// // // // // 	r := gin.Default()
// // // // // 	h := db.Init(dbUrl)

// // // // // 	populateDB(h)

// // // // // 	pizzashop.RegisterRoutes(r, h)

// // // // // 	r.GET("/", func(c *gin.Context) {
// // // // // 		c.JSON(200, gin.H{
// // // // // 			"port":  port,
// // // // // 			"dbUrl": dbUrl,
// // // // // 		})
// // // // // 	})

// // // // // 	r.Run()
// // // // // }

// // // // // func populateDB(db *gorm.DB) {
// // // // // 	var toppings []models.Topping
// // // // // 	var pizzas []models.Pizza
// // // // // 	var sizes []models.Size
// // // // // 	var crusts []models.Crust

// // // // // 	getData("F:/Onito_Assg/pkg/common/db/data/crusts.json", &crusts)
// // // // // 	getData("F:/Onito_Assg/pkg/common/db/data/pizzas.json", &pizzas)
// // // // // 	getData("F:/Onito_Assg/pkg/common/db/data/sizes.json", &sizes)
// // // // // 	getData("F:/Onito_Assg/pkg/common/db/data/toppings.json", &toppings)

// // // // // 	for _, topping := range toppings {
// // // // // 		db.Create(&topping)
// // // // // 	}

// // // // // 	for _, pizzas := range pizzas {
// // // // // 		db.Create(&pizzas)
// // // // // 	}

// // // // // 	for _, size := range sizes {
// // // // // 		db.Create(&size)
// // // // // 	}

// // // // // 	for _, crust := range crusts {
// // // // // 		db.Create(&crust)
// // // // // 	}
// // // // // }

// // // // // func getData(fileName string, v interface{}) {
// // // // // 	file, _ := os.Open(fileName)
// // // // // 	defer file.Close()
// // // // // 	byteValue, _ := ioutil.ReadAll(file)
// // // // // 	json.Unmarshal(byteValue, v)
// // // // // }

// // // package main

// // // import "fmt"

// // // func main() {
// // // 	a := map[string]int{"P1": 3, "P2": 2, "P3": 3, "P4": 4, "P5": 5} //tracks room assignments to each patients

// // // 	// one patient is assigned two rooms
// // // 	// one patient

// // // 	//step-1 : {"P1": 3, "P2": 2, "P3": 3, "P4":4, "P5": 5}
// // // 	//step-2 : {"P1": 3, "P2": 4, "P3": 3, "P4":4, "P5": 5}
// // // 	//step-3 : {"P1": 3, "P2": 4, "P3": 2, "P4":4, "P5": 5}
// // // 	//step-4 : {"P1": 3, "P2": 4, "P3": 2, "P4":5, "P5": 5}
// // // 	//step-5 : {"P1": 3, "P2": 4, "P3": 2, "P4":5, "P5": 1}

// // // 	v := []string{} //used to store discrepancies

// // // 	newMp := map[int]string{} //to keep track whether the patients are assigned to same room, if so arise a discrepancies.

// // // 	for key, element := range a {
// // // 		if newMp[element] != "" {
// // // 			v = append(v, fmt.Sprintf("%s and %s are alloted same room %d", newMp[element], key, element))
// // // 		} else {
// // // 			newMp[element] = key
// // // 		}
// // // 	}

// // // 	fmt.Println(v)
// // // }

// // // // package main

// // // // import (
// // // // 	"fmt"

// // // // 	"gonum.org/v1/gonum/graph/simple"
// // // // )

// // // // // NodeValue is a struct to store node name and value.
// // // // type NodeValue struct {
// // // // 	Name  string
// // // // 	Value int
// // // // }

// // // // // CustomNode is a custom node type that implements graph.Node.
// // // // type CustomNode struct {
// // // // 	*NodeValue
// // // // }

// // // // func (n *CustomNode) ID() int64 {
// // // // 	return int64(n.Value) // Use a unique ID for each node
// // // // }

// // // // func main() {
// // // // 	// Create a new directed graph
// // // // 	g := simple.NewDirectedGraph()

// // // // 	// Create custom nodes with values
// // // // 	nodes := map[string]*CustomNode{
// // // // 		"P1": {NodeValue: &NodeValue{Name: "P1", Value: 1}},
// // // // 		"P2": {NodeValue: &NodeValue{Name: "P2", Value: 2}},
// // // // 		"P3": {NodeValue: &NodeValue{Name: "P3", Value: 3}},
// // // // 		"P4": {NodeValue: &NodeValue{Name: "P4", Value: 4}},
// // // // 		"P5": {NodeValue: &NodeValue{Name: "P5", Value: 5}},
// // // // 	}

// // // // 	a := map[string]int{"P1": 3, "P2": 2, "P3": 3, "P4": 4, "P5": 5}

// // // // 	// Add nodes to the graph
// // // // 	for _, node := range nodes {
// // // // 		g.AddNode(node)
// // // // 	}

// // // // 	// Add edges based on the given information
// // // // 	// edges := []struct {
// // // // 	// 	from string
// // // // 	// 	to   string
// // // // 	// }{
// // // // 	// 	{"P1", "P3"},
// // // // 	// 	{"P3", "P5"},
// // // // 	// 	{"P5", "P2"},
// // // // 	// 	{"P2", "P4"},
// // // // 	// }

// // // // 	// for _, edge := range edges {
// // // // 	// 	fromNode := nodes[edge.from]
// // // // 	// 	toNode := nodes[edge.to]
// // // // 	// 	g.SetEdge(g.NewEdge(fromNode, toNode))
// // // // 	// }

// // // // 	// // Print the graph information
// // // // 	// nodesList := g.Nodes()
// // // // 	// edgesList := g.Edges()

// // // // 	// fmt.Println("Nodes:")
// // // // 	// for nodesList.Next() {
// // // // 	// 	node := nodesList.Node().(*CustomNode)
// // // // 	// 	fmt.Printf("%s (Value: %d)\n", node.Name, node.Value)
// // // // 	// }

// // // // 	// fmt.Println("Edges:")
// // // // 	// for edgesList.Next() {
// // // // 	// 	edge := edgesList.Edge()
// // // // 	// 	fmt.Printf("%s -> %s\n", edge.From().(*CustomNode).Name, edge.To().(*CustomNode).Name)
// // // // 	// }
// // // // }

// // // Billing example
// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"os"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/johnfercher/maroto/pkg/color"
// 	"github.com/johnfercher/maroto/pkg/consts"
// 	"github.com/johnfercher/maroto/pkg/pdf"
// 	"github.com/johnfercher/maroto/pkg/props"
// )

// //generating pdf manually

// func generatePDFContent(m pdf.Maroto) ([]byte, error) {
// 	var pdfContent []byte
// 	err := m.OutputFileAndClose("temp.pdf") // Generate a temporary PDF file
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Read the temporary file and get the PDF content
// 	pdfContent, err = ioutil.ReadFile("temp.pdf")
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Delete the temporary file
// 	err = os.Remove("temp.pdf")
// 	if err != nil {
// 		return nil, err
// 	}

// 	return pdfContent, nil
// }

// func generatePDF(c *gin.Context) {
// 	begin := time.Now()

// 	darkGrayColor := getDarkGrayColor()
// 	grayColor := getGrayColor()
// 	whiteColor := color.NewWhite()
// 	header := getHeader()
// 	contents := getContents()

// 	m := pdf.NewMaroto(consts.Portrait, consts.A4)
// 	m.SetPageMargins(10, 15, 10)

// 	m.RegisterHeader(func() {
// 		m.Row(20, func() {
// 			m.Col(3, func() {
// 				m.Text("AnyCompany Name Inc. 851 Any Street Name, Suite 120, Any City, CA 45123.", props.Text{
// 					Size:        8,
// 					Align:       consts.Right,
// 					Extrapolate: false,
// 				})
// 				m.Text("Tel: 55 024 12345-1234", props.Text{
// 					Top:   12,
// 					Style: consts.BoldItalic,
// 					Size:  8,
// 					Align: consts.Right,
// 				})
// 				m.Text("www.mycompany.com", props.Text{
// 					Top:   15,
// 					Style: consts.BoldItalic,
// 					Size:  8,
// 					Align: consts.Right,
// 				})
// 			})
// 		})
// 	})

// 	m.RegisterFooter(func() {
// 		m.Row(20, func() {
// 			m.Col(12, func() {
// 				m.Text("Tel: 55 024 12345-1234", props.Text{
// 					Top:   13,
// 					Style: consts.BoldItalic,
// 					Size:  8,
// 					Align: consts.Left,
// 				})
// 				m.Text("www.mycompany.com", props.Text{
// 					Top:   16,
// 					Style: consts.BoldItalic,
// 					Size:  8,
// 					Align: consts.Left,
// 				})
// 			})
// 		})
// 	})

// 	m.Row(10, func() {
// 		m.Col(12, func() {
// 			m.Text("Invoice ABC123456789", props.Text{
// 				Top:   3,
// 				Style: consts.Bold,
// 				Align: consts.Center,
// 			})
// 		})
// 	})

// 	m.SetBackgroundColor(darkGrayColor)

// 	m.Row(7, func() {
// 		m.Col(3, func() {
// 			m.Text("Transactions", props.Text{
// 				Top:   1.5,
// 				Size:  9,
// 				Style: consts.Bold,
// 				Align: consts.Center,
// 			})
// 		})
// 		m.ColSpace(9)
// 	})

// 	m.SetBackgroundColor(whiteColor)

// 	m.TableList(header, contents, props.TableList{
// 		HeaderProp: props.TableListContent{
// 			Size:      9,
// 			GridSizes: []uint{3, 4, 2, 3},
// 		},
// 		ContentProp: props.TableListContent{
// 			Size:      8,
// 			GridSizes: []uint{3, 4, 2, 3},
// 		},
// 		Align:                consts.Center,
// 		AlternatedBackground: &grayColor,
// 		HeaderContentSpace:   1,
// 		Line:                 false,
// 	})

// 	m.Row(20, func() {
// 		m.ColSpace(7)
// 		m.Col(2, func() {
// 			m.Text("Total:", props.Text{
// 				Top:   5,
// 				Style: consts.Bold,
// 				Size:  8,
// 				Align: consts.Right,
// 			})
// 		})
// 		m.Col(3, func() {
// 			m.Text("R$ 2.567,00", props.Text{
// 				Top:   5,
// 				Style: consts.Bold,
// 				Size:  8,
// 				Align: consts.Center,
// 			})
// 		})
// 	})

// 	m.Row(15, func() {
// 		m.Col(6, func() {
// 			_ = m.Barcode("5123.151231.512314.1251251.123215", props.Barcode{
// 				Percent: 0,
// 				Proportion: props.Proportion{
// 					Width:  20,
// 					Height: 2,
// 				},
// 			})
// 			m.Text("5123.151231.512314.1251251.123215", props.Text{
// 				Top:    12,
// 				Family: "",
// 				Style:  consts.Bold,
// 				Size:   9,
// 				Align:  consts.Center,
// 			})
// 		})
// 		m.ColSpace(6)
// 	})

// 	// Generate the PDF content
// 	pdfContent, err := generatePDFContent(m)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate PDF"})
// 		return
// 	}

// 	// Set the PDF content type and send it as response
// 	c.Data(http.StatusOK, "application/pdf", pdfContent)

// 	end := time.Now()
// 	fmt.Println(end.Sub(begin))
// }

// func main() {
// 	r := gin.Default()
// 	r.GET("/generate-pdf", generatePDF)
// 	// Run the server
// 	r.Run(":8080")
// }

// func getHeader() []string {
// 	return []string{"", "Product", "Quantity", "Price"}
// }

// func getContents() [][]string {
// 	return [][]string{
// 		{"", "Swamp", "12", "R$ 4,00"},
// 		{"", "Sorin, A Planeswalker", "4", "R$ 90,00"},
// 		{"", "Tassa", "4", "R$ 30,00"},
// 		{"", "Skinrender", "4", "R$ 9,00"},
// 		{"", "Island", "12", "R$ 4,00"},
// 		{"", "Mountain", "12", "R$ 4,00"},
// 		{"", "Plain", "12", "R$ 4,00"},
// 		{"", "Black Lotus", "1", "R$ 1.000,00"},
// 		{"", "Time Walk", "1", "R$ 1.000,00"},
// 		{"", "Emberclave", "4", "R$ 44,00"},
// 		{"", "Anax", "4", "R$ 32,00"},
// 		{"", "Murderous Rider", "4", "R$ 22,00"},
// 		{"", "Gray Merchant of Asphodel", "4", "R$ 2,00"},
// 		{"", "Ajani's Pridemate", "4", "R$ 2,00"},
// 		{"", "Renan, Chatuba", "4", "R$ 19,00"},
// 		{"", "Tymarett", "4", "R$ 13,00"},
// 		{"", "Doom Blade", "4", "R$ 5,00"},
// 		{"", "Dark Lord", "3", "R$ 7,00"},
// 		{"", "Memory of Thanatos", "3", "R$ 32,00"},
// 		{"", "Poring", "4", "R$ 1,00"},
// 		{"", "Deviling", "4", "R$ 99,00"},
// 		{"", "Seiya", "4", "R$ 45,00"},
// 		{"", "Harry Potter", "4", "R$ 62,00"},
// 		{"", "Goku", "4", "R$ 77,00"},
// 		{"", "Phreoni", "4", "R$ 22,00"},
// 		{"", "Katheryn High Wizard", "4", "R$ 25,00"},
// 		{"", "Lord Seyren", "4", "R$ 55,00"},
// 	}
// }

// func getDarkGrayColor() color.Color {
// 	return color.Color{
// 		Red:   55,
// 		Green: 55,
// 		Blue:  55,
// 	}
// }

// func getGrayColor() color.Color {
// 	return color.Color{
// 		Red:   200,
// 		Green: 200,
// 		Blue:  200,
// 	}
// }

// func getBlueColor() color.Color {
// 	return color.Color{
// 		Red:   10,
// 		Green: 10,
// 		Blue:  150,
// 	}
// }

// func getRedColor() color.Color {
// 	return color.Color{
// 		Red:   150,
// 		Green: 10,
// 		Blue:  10,
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"

// 	"github.com/benoitkugler/go-weasyprint/pdf"
// 	"github.com/benoitkugler/webrender/backend"
// 	"github.com/benoitkugler/webrender/html/document"
// 	"github.com/benoitkugler/webrender/html/tree"
// 	"github.com/benoitkugler/webrender/text"
// 	"github.com/benoitkugler/webrender/utils"
// )

// func HtmlToPdf(target io.Writer, htmlContent utils.ContentInput, fontConfig *text.FontConfiguration) error {
// 	return HtmlToPdfOptions(target, htmlContent, "", nil, "", nil, false, fontConfig, 1, nil)
// }

// func HtmlToPdfOptions(target io.Writer, htmlContent utils.ContentInput, baseUrl string, urlFetcher utils.UrlFetcher,
// 	mediaType string, stylesheets []tree.CSS, presentationalHints bool, fontConfig *text.FontConfiguration, zoom float64, attachments []backend.Attachment) error {
// 	parsedHtml, err := tree.NewHTML(htmlContent, baseUrl, urlFetcher, mediaType)
// 	if err != nil {
// 		return err
// 	}
// 	doc := document.Render(parsedHtml, stylesheets, presentationalHints, fontConfig)
// 	output := pdf.NewOutput()
// 	doc.Write(output, utils.Fl(zoom), attachments)
// 	pdfDoc := output.Finalize()
// 	return pdfDoc.Write(target, nil)
// }

// func main() {
// 	outputFile, err := os.Create("output.pdf")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer outputFile.Close()

// 	input := `
// 	<style>
// 		@page { size: 300px }
// 		svg { display: block }
// 	</style>
// 		<svg viewBox="-10 -10 120 120">
// 		<mask id="myMask">
// 			<!-- Everything under a white pixel will be visible -->
// 			<rect x="0" y="0" width="100" height="100" fill="white" />

// 			<!-- Everything under a black pixel will be invisible -->
// 			<path d="M10,35 A20,20,0,0,1,50,35 A20,20,0,0,1,90,35 Q90,65,50,95 Q10,65,10,35 Z" fill="black" />
// 		</mask>

// 		<polygon points="-10,110 110,110 110,-10" fill="orange" />

// 		<!-- with this mask applied, we "punch" a heart shape hole into the circle -->
// 		<circle cx="50" cy="50" r="50" mask="url(#myMask)" />
// 	</svg>
// 	`
// 	// Convert HTML to PDF

// 	err = HtmlToPdf(outputFile, utils.InputString(input), &text.FontConfiguration{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 		fmt.Println("PDF generation complete. Check output.pdf")
// 	}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"strings"

// 	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
// )

// func main() {
// 	// Create a new PDF generator
// 	pdfg, err := wkhtmltopdf.NewPDFGenerator()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Define your HTML content with embedded CSS
// 	htmlStr := `<html>
//     <head>
//         <style>
//             /* Your CSS styles here */
//             body {
//                 background-color: #0000FF;
//                 font-family: Arial, sans-serif;
//             }
//             h1 {
//                 color: red;
//             }
//         </style>
//     </head>
//     <body>
//         <h1>This is an HTML from PDF to test color</h1>
//         <img src="http://api.qrserver.com/v1/create-qr-code/?data=HelloWorld" alt="img" height="42" width="42" />
//         <!-- Your content here -->
//     </body>
// </html>`

// 	// Add the HTML content to the PDF generator
// 	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(htmlStr)))

// 	// Create the PDF
// 	err = pdfg.Create()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Save the PDF to a file
// 	err = pdfg.WriteFile("ayush.pdf")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Done")
// }

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/piang", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "poang",
		})
	})

	router.GET("/generate-pdf", GeneratePDF)

	// Start the HTTP server
	err := router.Run(":3000")
	if err != nil {
		log.Fatal(err)
	}
}

func GeneratePDF(c *gin.Context) {
	// Create a context for the browser
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Navigate to the HTML file using the file:// protocol
	if err := chromedp.Run(ctx, chromedp.Navigate("F:/Onito_Assg/cmd/sample_data/sample.html")); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to navigate to HTML file"})
		return
	}

	var htmlContent string
	if err := chromedp.Run(ctx, chromedp.Evaluate(`document.documentElement.outerHTML`, &htmlContent)); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get HTML content"})
		return
	}

	// Generate the PDF from the rendered HTML
	var pdfData []byte
	if err := chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
		var err error
		pdfData, _, err = page.PrintToPDF().Do(ctx)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate PDF"})
			return err
		}
		return nil
	})); err != nil {
		fmt.Println(err)

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to run PDF generation"})
		return
	}

	// Set the PDF data as the response content
	c.Data(http.StatusOK, "application/pdf", pdfData)
}
