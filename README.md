# gorchitect

A tool to visualize dependencies in a golang project.
This tool will help you to read the source code as you can quickly see the overall dependencies.
*Right now, only Windows is supported.*

# Installation

go get github.com/keikohi/gorchitect


# Usage 

(Windows)
``` cmd
gorchitect.exe -p "Absolute path of the golang project
```

Results of the analysis are output to the `result.dot` under this project.

The results can be visualized in [GraphvizOnline][link1]

[link1]:https://dreampuf.github.io/GraphvizOnline/


# Sample

The following figure shows the result of visualizing this project.
![graphviz](https://user-images.githubusercontent.com/11077015/95887279-9e8c3900-0dba-11eb-8f57-a9f08f43f398.png)
