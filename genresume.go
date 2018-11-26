package main

import (
	"github.com/burntsushi/toml"
	"github.com/jpfairbanks/structdoc/resume"
	"log"
	"os"
)

type School = resume.School
type Experience = resume.Experience
type Funding = resume.Funding
type Education = resume.Education
type Section = resume.Section
type ListSection = resume.ListSection
type Mentoring = resume.Mentoring
type Category = resume.Category
type Activity = resume.Activity
type Job = resume.Job
type Dates = resume.Dates
type Grant = resume.Grant
type Student = resume.Student
type ActivitySection = resume.ActivitySection
type Document = resume.Document

func NewResume() Document {
	edsec := []School{
		School{
			Name:         "Georgia Institute of Technology",
			Major:        "Ph.D Computational Science and Engineering",
			Advisor:      "Professor David A. Bader",
			Dissertation: "Graph Analysis Combining Numerical Statistical and Streaming Techniques",
			Qualifier:    "Computational Data Analysis (ML) and High Performance Computing (HPC)",
			Committee:    "Rich Vuduc, Haesun Park, Polo Chau, Dana Randall, G. Sanders (LLNL)",
			Address:      "Atlanta, GA",
			Dates:        "2012 -- 2016",
			Details:      "Research Assistant 2012, Teaching Assistant 2016"},
		School{
			Name:  "University of Florida",
			Major: "B.S. Mathematics",
			// Advisor: "Vince Vatter",
			Address: "Gainesville, FL",
			Details: "Summa cum laude",
			Thesis:  `A Ramsey Theorem for Indecomposable Matchings`,
			//\href{http://www.combinatorics.org/ojs/index.php/eljc/article/view/v18i1p227}{Abstract}
			Dates: "2009 -- 2012"}}

	experience := []Job{Job{
		Name:    `Georgia Tech Research Institute (GTRI)`,
		Address: `Atlanta, GA`,
		Title:   "Research Engineer II",
		Dates:   "May 2016 - Present",
		Details: []string{
			"Conduct research into high performance data analysis algorithms and applications",
			"Write grant proposals (see Funding)",
			"Manage federally funded research contracts",
			"Deliver applied research projects to sponsors such as source code, web applications, technical reports",
			"Mentor and advise students in connection to research projects"}},
		Job{
			Name:    "Ionic Security",
			Address: "Atlanta, GA",
			Title:   "Data Scientist",
			Dates:   "2015",
			Details: []string{"Developed data analytics software",
				"Designed a service oriented architecture for near real time analysis written in Go and Julia",
				"Leveraged time series and network database technologies including Heka, InfluxDB, RabbitMQ, and ElasticSearch"}},
		Job{
			Name:    "DOE -- Lawrence Livermore National Laboratory",
			Address: "Livermore, CA",
			Title:   "Institute for Scientific Computing Research Intern",
			Dates:   "2014",
			Details: []string{
				"Studied relationship between numerical accuracy of eigensolvers and solution quality of mincut graph partitioning",
				"Developed very fast approximate eigensolvers for large graphs",
				"Applied probabilistic reasoning to describe numerical computations",
				"Presented results at LLNL poster session"}},
		Job{
			Name:    "IDA -- Center for Computing Sciences",
			Address: "Bowie, MD",
			Title:   "Conducted research into Malware structure and similarities",
			Dates:   "2013",
			Details: []string{"Studied execution patterns of malicious programs",
				"Developed clustering and methods for understanding the structure of malicious programs with graph analytics",
				"Built a high performance distributed system for conducting these analyses with ZeroMQ communication"}},
	}
	students := []Student{Student{"Rohit Varkey", "MS CS Georgia Tech", 2018, "Google", "2016-2018"},
		Student{"Micah Halter", "BS CS Georgia Tech", 2019, "GT", "2016-Present"},
		Student{"Nate Knauf", "BS CS Georgia Tech", 2019, "GT", "2016"},
		Student{"Pushkar Godbole", "MS CSE Georgia Tech", 2016, "Yelp", "2015"}}

	mentoring := Mentoring{"Mentoring", "mentoring", students}
	research := ActivitySection{
		Name:  "Research",
		Label: "research",
		Categories: []Category{
			Category{
				Name: "Peer Reviewed Journal Articles",
				Activities: []Activity{
					Activity{Authors: "J. P. Fairbanks, R. Kannan, H. Park, D. A. Bader",
						Title: "Behavioral Clusters in Dynamic Graphs",
						Venue: "Parallel Computing Special Issue of Scientific Graph Analysis",
						Date:  "2015"},
					Activity{
						Authors: `J. P. Fairbanks`,
						URL:     "http://www.combinatorics.org/ojs/index.php/eljc/article/view/v18i1p227/pdf",
						Title:   "A Ramsey Theorem for Indecomposable Matchings",
						Venue:   "Electronic Journal of Combinatorics, Vol 18(1)",
						Date:    `Dec 2011`}}},
			Category{Name: "Peer Reviewed Conference Publications",
				Activities: []Activity{
					Activity{
						Authors: "N. Campbell, T. Goodyear, W. Messer, E. Stuart, J. P. Fairbanks",
						URL:     "http://ieee-hst.org/october23.htm",
						Title:   "Digital Witness: Remote Methods for Volunteering Digital Evidence on Mobile Devices",
						Venue:   "IEEE Technologies for Homeland Security",
						Date:    "Oct 2018"},
					Activity{
						Authors: "R. Varkey Thankachan, B. P. Swenson, J. P. Fairbanks",
						Title:   "Performance Effects of Backing Data Stores in Community Detection Algorithms",
						URL:     "http://ieee-hpec.org/copy/agendatext.html",
						Venue:   "IEEE High Performance Extreme Computing",
						Date:    "Sep 2018"},
					Activity{
						Authors: "N. Fitch, N. Knauf, J. P. Fairbanks, E. Briscoe",
						URL:     "http://snap.stanford.edu/mis2/files/MIS2_paper_17.pdf",
						Title:   "Credibility Assessment in the News: Do we need to read?",
						Venue:   "ACM WSDM MIS2",
						Date:    "Feb 2018"},
					Activity{
						Authors: "R. Varkey Thankachan, E. Hein, B. P. Swenson, J. P. Fairbanks",
						Title:   "Integrating Productivity-Oriented Programming Languages with High-Performance Data Structures",
						URL:     "https://ieeexplore.ieee.org/document/8091068/",
						Venue:   "IEEE High Performance Extreme Computing",
						Date:    "Sep 2017"},
					Activity{
						Authors: "J. P. Fairbanks, D. M. Ediger",
						URL:     "http://graphanalysis.org/GABB17_program_final.pdf",
						Title:   `Deriving Streaming Graph Algorithms from Static Definitions`,
						Venue:   "IEEE International Parallel and Distributed Processing Graph Algorithms Building Blocks",
						Date:    "2017"},
					Activity{
						Authors: "J. P. Fairbanks, D. A. Bader, and G. D. Sanders",
						Title:   "Graph Partitioning with Spectral Blends",
						Venue:   "Oxford Journal of Complex Networks",
						Date:    "Jan 2017"},
					Activity{
						Authors: "E. Nathan, G. Sanders, J. P. Fairbanks, V. Henson and D. Bader",
						Title:   "Graph Ranking Guarantees for Numerical Approximations to Katz Centrality",
						Venue:   `International Conference On Computational Science`,
						Date:    "2017"},
					Activity{
						Authors: "D. M. Ediger and J. P. Fairbanks",
						Title:   `Deriving Streaming Graph Algorithms from Static Definitions.`,
						Venue:   `IEEE Parallel and Distributed Processing - Graph Algorithm Building Blocks`,
						Date:    "2017"},
					Activity{
						Authors: "A. Zakrzewska, E. Nathan, J. P. Fairbanks, D. A. Bader",
						Title:   `A local measure of community change in dynamic graphs.`,
						Venue:   "IEEE/ACM ASONAM",
						// Details: []string{"349-353"},
						Date: ""},
					Activity{
						Authors: "J. P. Fairbanks, A. Zakrzewska, D.A. Bader",
						URL:     "www.siam.org/meetings/ns16/",
						Title:   "Novel Stopping Criteria for Spectral Partitioning",
						Venue:   "SIAM Network Science",
						Date:    "Jul 2016"},
					Activity{
						Authors: "J. P. Fairbanks, D. Ediger, R. McColl, D.A. Bader, E. Gilbert",
						Title:   "A Statistical Framework for Analyzing Streaming Graphs",
						URL:     "http://stingergraph.com/data/uploads/papers/streaming-twitter-stats.pdf",
						Venue:   "IEEE/ACM ASONAM",
						Date:    "Aug 2013"},
				}},
			Category{Name: "Oral Presentations",
				Activities: []Activity{
					Activity{
						Title:   "Data Science and Graph Analytics with Julia",
						Details: []string{"Host: UF Data Science and Informatics"},
						Authors: "J. P. Fairbanks",
						Venue:   "University of Florida Informatics Institute",
						Date:    "Nov 2018"},
					Activity{
						Title:   "Solving Applied Graph Theory Problems in the JuliaGraphs ecosystem",
						Venue:   "MIT CSAIL Seminar",
						Date:    "2018",
						Details: []string{"Host: Alan Edelman, MIT Math/CSAIL"},
						Authors: "J. P. Fairbanks"},
					Activity{
						Title:   "Graph Interfaces: Bespoke Graphs for Every Occasion",
						Authors: "M. Besan\\c{c}on, J. P. Fairbanks",
						URL:     "https://youtu.be/OD-BSn4FZ2A",
						Venue:   "JuliaCon, London, UK",
						Date:    "2018"},
					Activity{
						Title:   "The JuliaGraphs Ecosystem: Move Fast and Don't Break Things",
						Authors: "J. P. Fairbanks",
						URL:     "https://youtu.be/OZuQoxTPoyM",
						Venue:   "JuliaCon, London, UK",
						Date:    "2018"},
					Activity{
						Title:   "Assessing Credibility in Global Media Networks",
						Authors: "J. P. Fairbanks, Human Language Technologies",
						Date:    "2017"},
					Activity{
						Title:   "Using Big Data to Predict and Analyze Cooperation and Conflict",
						Authors: "T. Frederick, C. Herlihy, J. P. Fairbanks",
						Venue:   "The Conflict Conference at UT-Austin",
						Date:    "2017"},
					Activity{
						Title:   "LightGraphs: Our Network, Our Story",
						Venue:   "JuliaCon, Berkeley, CA",
						Authors: "S. Bromberger, J. P. Fairbanks",
						URL:     "https://youtu.be/MFD-qmApXl8",
						Date:    "2017"}}},
			Category{Name: "Posters",
				Activities: []Activity{
					Activity{
						Title:   "QueryGarden: growing healthy applications in well prepared SQL",
						Authors: "J. P. Fairbanks",
						Venue:   "OHDSI Symposium",
						Date:    "2017"},
					Activity{
						Title:   "Implementing Real-Time Patient Level Predictions Using PLP Models",
						Authors: "C. S. Brown, J. D. Duke, , J. P. Fairbanks, C. Herlihy, K. Mukadam, J. Poovey, M. Rost",
						Venue:   "OHDSI Symposium",
						Date:    "2017"},
					Activity{
						Title: "Discovering Block Structure with Approximate Eigenvectors",
						Venue: "SIAM Computational Science and Engineering",
						Date:  "Mar 2015"},
					Activity{
						Title:   "Ramsey Theorem for Indecomposable Matchings",
						Venue:   "Graph Theory at Georgia Tech (GT@GT)",
						Authors: "",
						Date:    "2012"}}}}}
	oss := []string{`Core maintainer of \href{github.com/JuliaGraphs/LightGraphs.jl}{\emph{LightGraphs}} the most widely used Graph Algorithm Package in \href{julialang.org}{\emph{Julia}}.`,
		`Developer of \href{stingergraph.com}{\emph{STINGER}} the fastest streaming dynamic graph library for shared memory parallel computers.`}
	teaching := []Category{
		Category{
			"Professional Education",
			[]Activity{
				Activity{
					Date:  "Fall 2018",
					Title: "Programming for Data Science with Beverly Wright"}}},
		Category{
			"Georgia Tech Professional Education",
			[]Activity{
				Activity{
					Date:  "Spring 2017",
					Title: "Data Analytics Methodology with J. Poovey, D. Ediger, and M. Rost."},
				Activity{
					Date:  "Fall 2016",
					Title: "Big Data Analytics with J. Poovey, D. Ediger, and M. Rost."}}},
		Category{"Teaching Assistant at Georgia Tech",
			[]Activity{
				Activity{
					Title: "CSE 6643 Numerical Linear Algebra with Prof. Haesun Park",
					Date:  "Spring 2016"},
				Activity{
					Title: "CSE 6220 High Performance Computing with Prof. Srinivas Aluru",
					Date:  "Spring 2014"}}}}
	honors := Category{Name: "Honors, Awards, and Fellowships",
		Activities: []Activity{
			Activity{Title: "Presidential Fellowship for Graduate Study at Georgia Tech", Date: "2012 - 2016"},
			Activity{Title: "University Scholar at the University of Florida", Date: "2011 - 2012"},
			Activity{Title: "Kermit Sigmon Scholarship *for service to the mathematical community*", Date: "2012"},
			Activity{Title: "Tau Beta Pi, Engineering Honor Society, Georgia Tech Chapter", Date: "2015"},
			Activity{Title: "Phi Beta Kappa, University of Florida Chapter", Date: "2012"}}}

	leadership := Category{Name: "Leadership and Service",
		Activities: []Activity{
			Activity{
				Title: "JuliaCon Organizing Committee Vice Program Chair",
				Date:  "2018",
				Details: []string{"Organized the technical program of a 3 day international conference on the Julia programming language",
					"Ran Program Committee meetings to decide on accepted abstracts and presentations",
					"Led poster session preparations"}},
			Activity{
				Title:   "Tau Beta Pi Atlanta Alumni Chapter President",
				Date:    "2017",
				Details: []string{"Organized professional networking events for local Atlanta Area Engineers"}},
			Activity{
				Title: "Georgia Tech College of Computing Graduate Student Association VP for the School of CSE",
				Date:  "2015",
				Details: []string{"Represented department students to university administration committees on curriculum and funding",
					"Organized social and professional networking events for graduate students",
					"Chaired the organizing committee of HotCSE graduate research seminar providing early career presentation opportunities to graduate students"}},
			Activity{
				Title:   "Univ. Florida Pi Mu Epsilon Chapter President",
				Date:    "2011",
				Details: []string{"Organized a series of talks for the mathematics students at UF on diverse mathematical topics and skills incl. LaTeX, programming and technical communication in the field."}},
			Activity{
				Title: "Eagle Scout",
				Date:  "2009"}}}
	achievements := []Category{honors, leadership}

	skills := []string{
		"Programming languages (most familiar to least) Julia, Golang, Python, C, SQL, Bash, Matlab",
		"Computational Data Analysis (pandas, sklearn, Jupyter)",
		"Web development with Golang and Python (flask)",
		"Database Applications primarily with PostgreSQL and MongoDB",
		"Practical computing skills such as \texttt{*NIX}, git, make, \\LaTeX",
		"Continuous Integration/Deployment: Docker, DC/OS, Kubernetes",
		"Avid Linux User"}

	darpa := "DARPA"
	PI := "Principal Investigator"
	grants := []Grant{
		Grant{
			Role:      PI,
			Agency:    darpa,
			Title:     "Artificial Intelligence Exporation -- Automating Scientific Knowledge Extraction",
			Agreement: "Agreement No.~HR00111990008",
			Dates:     Dates{"2018", "2020"},
			Amount:    "≈1M"},
		Grant{
			Role:      PI,
			Agency:    "National Inst. of Justice",
			Title:     "Developing Novel Means of Evidence Collection",
			Agreement: "Grant Number 2016-MU-MU-K004",
			Dates:     Dates{"2016", "2018"},
			Amount:    "≈400K"},
		Grant{
			Role:   "Winner",
			Agency: "Office of the Director of National Intelligence",
			Title:  "XAMINE Challenge",
			Dates:  Dates{"2018", ""}},
		Grant{
			Role:   "Fellow",
			Agency: "American Society of Engineering Education",
			Dates:  Dates{"2013", "2016"},
			Title:  "National Defense Science and Engineering Fellowship"},
		Grant{
			ProjectID: "0D79670000",
			Role:      "Task Lead",
			Agency:    "Office of Naval Research",
			Title:     "Performance Estimation of Underwater MCM Operations",
			Amount:    "≈990K",
			Agreement: "Contract No. N00014-16-C-3041",
			//Aug Jul
			Dates: Dates{"2016", "2019"},
			//Jan
			Involved: Dates{"2018", "Present"}},
		Grant{
			Role:   "Task Lead",
			Agency: "GTRI Strategic Initiative",
			Title:  "Multi-source Anticipatory Intelligence",
			Dates:  Dates{"2016", "2019"}},
		Grant{
			ProjectID: "0D77150000",
			Role:      "Performer",
			Agency:    "Office of Naval Research",
			Amount:    "540K",
			Agreement: "Contract No. N00014-15-C-5172",
			Involved:  Dates{"2018", "2019"},
			Title:     "Automation for UxV-based Mine Countermeasures",
			Dates:     Dates{"2015", "2019"}},
		Grant{
			Role:   "Performer",
			Agency: "GTRI Strategic Initiative",
			Title:  "Healthy Wealthy Wise",
			Dates:  Dates{"2016", "2017"}}}
	doc := Document{
		Title:  "Resume",
		Author: "James Fairbanks",
		Name:   "James Fairbanks, PhD",
		Email:  "James.Fairbanks@gtri.gatech.edu",
		URL:    "jpfairbanks.com",
		GitHub: "jpfairbanks",
		Date:   "\\today",
		// Header:    header,
		Education: Education{Section{Name: "Education", Label: "Education"}, edsec},
		Experience: Experience{
			Section: Section{Name: "Work Experience", Label: "work-experience"},
			Items:   experience},
		Funding:      Funding{Name: "Funding", Label: "funding", Grants: grants},
		Teaching:     ActivitySection{"Teaching", "teaching", teaching},
		Mentoring:    mentoring,
		Research:     research,
		OpenSource:   ListSection{"Open Source", "open-source", oss},
		Achievements: ActivitySection{"Achievements", "achievements", achievements},
		Skills:       ListSection{"Selected Technical Skills", "skills", skills},
		Body:         "Hello Latex"}
	return doc
}

func main() {

	doc := NewResume()
	conffile, err := os.Create("cv.toml")
	if err != nil {
		log.Fatalf("Failed to create output toml file: %s", err)
	}
	enc := toml.NewEncoder(conffile)
	enc.Encode(doc)
}
