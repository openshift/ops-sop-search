import React, { Component } from 'react'
import { extend } from "lodash";
import { SearchkitManager,SearchkitProvider,
  SearchBox, Pagination,
  HitsStats, SortingSelector, ResetFilters, 
  GroupedSelectedFilters, Layout, TopBar, InputFilter,
  LayoutBody, LayoutResults, ActionBar, Hits, 
  NoHits, ActionBarRow, SideBar, TagFilterList } from 'searchkit'
import './index.css'
// import { Remarkable } from 'remarkable';
const host = "http://localhost:9200/sop/" //this is where you put the elastic url from the route you created
const searchkit = new SearchkitManager(host)
const asciidoctor = require('asciidoctor')()

// var md = new Remarkable();

const SOPItem = (props) => {
  const {bemBlocks, result} = props
  const source = extend({}, result._source, result.highlight)

  if (source.tags.includes("markdown")){

    return (
      <div className={bemBlocks.item().mix(bemBlocks.container("item"))} data-qa="hit">
        <div className={bemBlocks.item("details")}>
          <h2 className={bemBlocks.item("title")} dangerouslySetInnerHTML={{__html:source.name}}></h2>
          <ul className={bemBlocks.item("tags")}>
            <li>Tags: <TagFilterList field="tags.raw" values={source.tags} /></li>
            <li>Authors: <TagFilterList field="author.raw" values={source.author} /></li>
          </ul>
          <div className={bemBlocks.item("text")} dangerouslySetInnerHTML={{__html:source.content}}></div>
        </div>
      </div>
     )
  }
  return (
    <div className={bemBlocks.item().mix(bemBlocks.container("item"))} data-qa="hit">
      <div className={bemBlocks.item("details")}>
        <h2 className={bemBlocks.item("title")} dangerouslySetInnerHTML={{__html:source.name}}></h2>
        <ul className={bemBlocks.item("tags")}>
          <li>Tags: <TagFilterList field="tags.raw" values={source.tags} /></li>
          <li>Authors: <TagFilterList field="author.raw" values={source.author} /></li>
        </ul>
        <div className={bemBlocks.item("text")} dangerouslySetInnerHTML={{__html:asciidoctor.convert(source.content)}}></div>
      </div>
    </div>
   )
}

class App extends Component {
  render() {
    return (
      <SearchkitProvider searchkit={searchkit}>
      <Layout>
        <div className="top">
          <TopBar>
            <div className="my-logo">Openshift SOP Searcher</div>
            <SearchBox autofocus={true} searchOnChange={true} prefixQueryFields={["name", "content"]} queryFields={["author","name^10","tags","content^5"]}/>
          </TopBar>
        </div>
      <LayoutBody>

        <SideBar>
          <InputFilter id="name" searchThrottleTime={500} title="Document Name" placeholder="Search names" searchOnChange={true} prefixQueryFields={["name"]} queryFields={["name"]} />
          <InputFilter id="tags" searchThrottleTime={500} title="Tags" placeholder="Search tags" searchOnChange={true} prefixQueryFields={["tags"]} queryFields={["tags"]} />
          <InputFilter id="author" searchThrottleTime={500} title="Author" placeholder="Search authors" searchOnChange={true} prefixQueryFields={["author"]} queryFields={["author"]} />
        </SideBar>

        <LayoutResults>

          <ActionBar>
            <div className="body">
            <ActionBarRow>
                <HitsStats/>
                  <SortingSelector options={[
                    {label:"Relevance", field:"_score", order:"desc", defaultOption:true},
                    {label:"Last Updated", field:"lastUpdated", order:"desc"}
                  ]}/>  
            </ActionBarRow>
            </div>
            <ActionBarRow>
              <GroupedSelectedFilters/>
              <ResetFilters/>
            </ActionBarRow>

          </ActionBar>
          <div className="body">
            <Hits hitsPerPage={1} highlightFields={["name", "links", "content"]} sourceFilter={["name", "content", "tags", "author"]}
        mod="sk-hits-list" itemComponent={SOPItem}/>
            <NoHits suggestionsField={"name"}/>
          </div>
          <Pagination showNumbers={true}/>

        </LayoutResults>

      </LayoutBody>

      </Layout>
      </SearchkitProvider>
    );
  }
}
export default App;
