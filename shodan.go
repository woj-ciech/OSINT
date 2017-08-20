package main

import (
    "fmt"   // printing
    "gopkg.in/ns3777k/go-shodan.v1/shodan" // shodan api
    "github.com/PuerkitoBio/goquery"
    "log"
    "os" // os commands
    "bufio" //input output
    "strconv"
)

func bugcrowd() []string{
	doc, err := goquery.NewDocument("https://www.bugcrowd.com/bug-bounty-list/") //link to bugcrowd list
        if err != nil { // if cant connect
                log.Fatal(err)
        }
	
	s := make([]string, 455 ) //slice of strings

	doc.Find("td a").Each(func(index int, item *goquery.Selection) { //for every organization in "td a" (table)
	    linkTag := item.Text() // get text
	    s[index] = linkTag //put in to map
	    })
	
	return s // return slice of organizations
}

func intro() (){
	fmt.Println("Search multiple organization with Shodan.")
	fmt.Println("Prepare txt file with organizations names")
	fmt.Println("and pass it to script as argument")
	fmt.Println("----------------------------------------")
	fmt.Println("Example")
	fmt.Println("Hosts.txt includes:")
	fmt.Println("Sony")
	fmt.Println("Facebook")
	fmt.Println("Dropbox")
	fmt.Println("---------------------------------------")
	fmt.Println("Run ./shodan hosts.txt")
	fmt.Println("As output script makes directory with organization's name")
	fmt.Println("and writes response as txt file")
	fmt.Println("/Sony")
	fmt.Println("---xxx.xxx.xxx.xxx")
	fmt.Println("---xxx.xxx.xxx.xxx")
	fmt.Println("/Facebook")
	fmt.Println("---xxx.xxx.xxx.xxx")
	fmt.Println("---xxx.xxx.xxx.xxx")
}

func in_array(val int, array []int) (ok bool, i int) { //https://codereview.stackexchange.com/questions/60074/in-array-in-go
    for i = range array {
        if ok = array[i] == val; ok {
            return
        }
    }
    return
}

func func_unique(input []string) []string { //https://kylewbanks.com/blog/creating-unique-slices-in-go
	u := make([]string, 0, len(input))
	m := make(map[string]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}

	return u
	}

func read_file(input string) (orgs_string []string) { // https://stackoverflow.com/questions/8757389/reading-file-line-by-line-in-go
	//var orgs_string [400]string
	file, err := os.Open(input)
    if err != nil {
        log.Fatal(err)
		os.Exit(1)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        //fmt.Println(scanner.Text())
        orgs_string = append(orgs_string, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        intro() 
		os.Exit(1)
    }

    return orgs_string
}

/* function for infinite loop in user input;  when user type 111 - exit from loop 
it gets slice of strings (uniique) and return index and breaking point*/

func loop(query []string) (breaking bool, index int) {
		index = 0
		breaking = false
		var input int
	    fmt.Scan(&input) // get input
		if input == 111 {
			breaking = true 
			return breaking, index // return false (do not exit from main loop), return index = 0
		}else{
			index = input 
			breaking = false
			return breaking, index 
			}
	return breaking, index
}	


/*Printing loop, gets slice of unique organizations and map{org:value of appaerances}
Needs for counting results */
func print_loop(orgs []string, counter map[string]int)  {
	for index, value :=range orgs{
		fmt.Println(index,"-------->", value, "-", counter[value], "results for that organization") //get value from counter for every organizations in unique
		 
	}
	fmt.Println("Type 111 to go to next organization")
}
func main() {

	if len(os.Args) == 1 {
		intro() 
		os.Exit(1)
	}
	argsWithoutProg := os.Args[1]
	// multiple ports query, dont need to check if port is in array
	ports := "port:'80, 81, 443, 8000, 8001, 8008, 8080, 8083, 8443, 8834, 8888'"
	// ports := []int{80, 81, 443, 591, 2082, 2095, 2096, 3000, 8000, 8001, 8008, 8080, 8083, 8443, 8834, 8888, 55672}
	
	orgquery := read_file(argsWithoutProg) //
	

    client := shodan.NewClient(nil, "YOUR_API_KEY") //") //connect with shodan

    //fmt.Println(client.GetAPIInfo()) //Api info

    for i:= range orgquery{ // for every item in slice
    	organtion_query := "org:" + orgquery[i] + " " + ports //make string "org: Company ports: ....."
    	fmt.Println("Query:", organtion_query) // print query

	    a := &shodan.HostQueryOptions{Query:organtion_query} //first query "org: Company port: ....."

	    query, err := client.GetHostsForQuery(a) /*get host from previous query and return 
	    type HostMatch struct {
    	Total   int                 `json:"total"`
    	Facets  map[string][]*Facet `json:"facets"`
    	Matches []*HostData         `json:"matches"` }*/

	    fmt.Println("Error:", err) //print error
	    fmt.Println("Founded: ", query.Total, "total result(s)") // print amount of results for everything
	    if query.Total == 0 { // debugging 
						fmt.Println("Nothing was found. Going to next one\n")
						continue
						
					} 
	    organization := make([]string, len(query.Matches)) //slice for organizations

	    for i:= range query.Matches { //for every match in query
	    	organization[i] = query.Matches[i].Organization //get organization
	    	// now organization is [Company1, Company2 ...] with repetives
	    	}

		unique := func_unique(organization) // get uniqie from organization
		counts := make(map[string]int) //map for counting orgs 
			for _, helperorg := range organization{ // for every value in organization
				counts[helperorg]++ // add to map and ++
			}
		fmt.Println("Organizations:")
		//print_loop(unique)
		

		print_loop(unique, counts)
		i = 0
		// starbuck. sony
		

			for{ // for {} means infinite loop
				breaking_point, index := loop(unique) // 
				if breaking_point == false {
					if index >= len(unique){
						fmt.Println("Wrong choice, try again")
						print_loop(unique, counts)
						continue
					} 
					fmt.Println("Company: ", unique[index]) // print founded organizations
					uniquery := "org:'" + string(unique[index]) + "' " + ports // make query for next request {org: i from unique organizations}
					fmt.Println("Query:", uniquery)
					c := &shodan.HostQueryOptions{Query:uniquery} // connect and query
					query1, _ := client.GetHostsForQuery(c) // make request
					//fmt.Println(query1.Total)


					//fmt.Println(query1.Total) // print total results from unique organizations for EVERY COMPANY!!
					for j:= range query1.Matches{ // for every matches from query1
						path := unique[index] + "/" + query1.Matches[j].IP // company/111.111.111.111
						fmt.Println("Making file in ", path)
						os.Mkdir(unique[index], os.FileMode(0666)) // make directory with 
						f, _ := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666) // read and write for everyone
						w := bufio.NewWriter(f) // make new writer
						s := strconv.Itoa(query1.Matches[j].Port)
						fmt.Fprintf(w, "Port:")
						fmt.Fprintf(w,  s)
						fmt.Fprintf(w, "\n")
						//fmt.Fprintf(w, "\n")
						if len(query1.Matches[j].Hostnames) > 0 { //if there is hostname related to match
							for a:= range query1.Matches[j].Hostnames{ // for every founded hostname
								fmt.Fprintln(w, query1.Matches[j].Hostnames[a]) //add to file
								//fmt.Fprintf(w, "\n")
						}
						}
						fmt.Fprintf(w, query1.Matches[j].Data) //save data (http response) to new writer

						w.Flush() // flussssssssssssssh				
						i++
				}
				fmt.Println("\n")
				print_loop(unique, counts)
		}else{
					break
					}
			}		
		}
	}
