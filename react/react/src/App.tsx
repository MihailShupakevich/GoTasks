import { useMemo, useState } from 'react';
// import  Counter  from "./components/Counter";
// import ClassCounter from "./components/ClassCounter"
// import NewComponent from "./components/NewComponent"
// import NewRefComponent from "./components/NewRefComponent"
// import NewC from "./components/NewC"
import  './styles/App.css'
import PostList from  "./components/PostList";
import PostForm from  "./components/PostForm";
import MySelect from  "./components/UI/select/MySelect";
import MyInput from  "./components/UI/input/MyInput";


function App() {

  const [posts, setPosts] = useState([
    {id:1,title:'Javascript',body:'ABS'},
    {id:2,title:'Javascript2',body:'bb'},
    {id:3,title:'Javascript3',body:'sS'},
    {id:4,title:'Javascript4',body:'ccBS'},
    {id:5,title:'Javascript5',body:'A23442S'}])
const [selectedSort,setSelectedSort] = useState('')
const [searchQuery, setSearchQuery] = useState('')


const sortedPosts = useMemo( ()=> {
  if(selectedSort){
    return [...posts].sort((a,b)=>a[selectedSort].localeCompare(b[selectedSort]))
  }
  return posts;
   
},[selectedSort, posts]);

const createPost = (newPost) => {
  setPosts([...posts,newPost])
}
const removePost = (post) => {
  setPosts(posts.filter(p => p.id !== post.id))
}
const sortPosts = (sort) => {
  setSelectedSort(sort);

  
}
  return (

  <div className="App" >
    <PostForm  create={createPost} />
    <hr style={{margin:'15px'}} />
    <div>
      <MyInput placeholder='Poisk...' 
      value={searchQuery}
      onChange={e=>setSearchQuery(e.target.value) }
      />
    <MySelect
    value= {selectedSort}
    onChange = {sortPosts}
    defaultValue='Cортировка'
    options={[
    {value:'title', name:"По названию"},
    {value:'body', name:"По описанию"}
    ]}
    />
    </div>
    {posts.length !== 0
    ?(<PostList remove={removePost} posts={sortedPosts} title='Hey!' />)
    :(<h1 style={{textAlign:'center'}}>Пoсты не найдены!!! </h1>)}
    </div>
  );
}
//Катя поставь аву
export default App;
