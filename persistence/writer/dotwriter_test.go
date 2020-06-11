package writer

import (
	"testing"
	"time"
)

func TestWriter(t *testing.T) {

	// // dir := "D:\\dev\\go\\golang\\go"
	// dir := "D:\\dev\\go\\voxelize"
	// filepaths := findGoFiles(dir)
	// nf := newNodeFactory(filepaths)
	// nf.newProjectNodes()
	// gf := newGraphFactory(nf.rootNode)
	// rootGraph := gf.create()

	// fp, err := os.Create("./test.dot")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// writer := Dotwriter{w: fp, root: rootGraph}
	// writer.Write(rootGraph)

}

func TestAll(t *testing.T) {

	// // dir := "D:\\dev\\go\\voxelize"
	// // dir := "D:\\dev\\go\\golang\\go"
	// // dir := "D:\\dev\\go\\docker-ce-master\\docker-ce-master"
	// // dir := "D:\\dev\\go\\docker-ce-master\\docker-ce-master\\components"
	// // dir := "D:\\dev\\go\\gin-master\\gin-master"
	// // dir := "D:\\dev\\go\\anko-master"
	// // dir := "D:\\dev\\go\\redigo"
	// dir := "D:\\dev\\go\\goviz-master"
	// // dir := "D:\\dev\\go\\goviz-master\\goviz-master"
	// // dir := "D:\\dev\\go\\anko-master"
	// // dir := "D:\\dev\\go\\serf-master"
	// filepaths := findGoFiles(dir)
	// nf := newNodeFactory(filepaths)
	// nf.newProjectNodes()
	// gf := newGraphFactory(nf.rootNode)
	// rootGraph := gf.create()

	// da := DependencyAnalyzer{filepaths: filepaths}
	// codes := da.Analyse()
	// deps := codes.dependency()

	// fp, err := os.Create("./test.dot")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// writer := Dotwriter{w: fp, root: rootGraph, deps: deps.mergeDuplicate()}
	// writer.Write(rootGraph)

	// // da := DependencyAnalyzer{filepaths: filepaths}
	// // gofiles := da.Analyse()
	// // deps := gofiles.dependency()
	// // fmt.Println(deps)
}

type tm struct {
	name     string
	duration time.Duration
}

func TestFocus(t *testing.T) {
	// // dir := "D:\\dev\\go\\golang\\go"

	// times := make([]tm, 0, 20)
	// start := time.Now()
	// // dir := "D:\\dev\\go\\docker-ce-master\\docker-ce-master\\components"
	// dir := "D:\\dev\\go\\anko-master"
	// filepaths := findGoFiles(dir)
	// times = append(times, tm{"findGoFiles", time.Since(start)})
	// nf := newNodeFactory(filepaths)
	// times = append(times, tm{"newNodeFactory", time.Since(start)})
	// nf.newProjectNodes()
	// times = append(times, tm{"newProjectNodes", time.Since(start)})
	// gf := newGraphFactory(nf.rootNode)
	// times = append(times, tm{"newGraphFactory", time.Since(start)})
	// // rootDotGraph := gf.create()

	// da := DependencyAnalyzer{filepaths: filepaths}
	// gofiles := da.Analyse()
	// times = append(times, tm{"Analyse", time.Since(start)})
	// deps := gofiles.dependency()
	// times = append(times, tm{"dependency", time.Since(start)})
	// focuseddeps, chain := deps.focusPackage("main")
	// times = append(times, tm{"focusPackage", time.Since(start)})
	// focusedDotGraph := gf.focusPackage(chain)
	// times = append(times, tm{"focusPackage", time.Since(start)})

	// fp, err := os.Create("./test.dot")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// writer := Dotwriter{w: fp, root: focusedDotGraph, deps: focuseddeps.mergeDuplicate()}
	// times = append(times, tm{"Dotwriter", time.Since(start)})
	// writer.Write(focusedDotGraph)
	// times = append(times, tm{"Write", time.Since(start)})

	// tmp := 0.0
	// for _, v := range times {
	// 	fmt.Printf("%2.8f / %2.8f sec <= %s\n", v.duration.Seconds()-tmp, v.duration.Seconds(), v.name)
	// 	tmp = v.duration.Seconds()
	// }

}
