import { useRef, useState } from "react";
// import  Counter  from "./components/Counter";
// import ClassCounter from "./components/ClassCounter"
// import NewComponent from "./components/NewComponent"
// import NewRefComponent from "./components/NewRefComponent"
// import NewC from "./components/NewC"
import  './styles/App.css'
import PostList from  "./components/PostList";
import MyButton from  "./components/UI/button/MyButton";
import MyInput from  "./components/UI/input/MyInput";
function App() {

const [posts, setPosts] = useState([
  {id:1,title:'Javascript',body:'ABS'},
  {id:2,title:'Javascript',body:'bb'},
  {id:3,title:'Javascript',body:'sS'},
  {id:4,title:'Javascript',body:'ccBS'},
  {id:5,title:'Javascript',body:'A23442S'}])

 const [post,setPost] =useState( {title:'', body:''})

  const addNewPost = (e)=>{
    e.preventDefault();
    setPosts([...posts,{...post,id:Date.now(),}])
    setPost( {title:'', body:''})
    
  }

  return (

  <div className="App" >
    <form >
      <MyInput 
      onChange = {e=> setPost({...post, title: e.target.value})}
      value={post.title}
       type="text" 
       placeholder="postName"/>
      <MyInput 
      type="text"
       placeholder="Описание поста" 
       onChange = {e=> setPost({...post,body: e.target.value})}
       value={post.body}/>
      <MyButton onClick={addNewPost}>Cоздать пост </MyButton>
    </form>
    <PostList posts={posts} title='Hey!' />
    </div>
  );
}
//Катя поставь аву
export default App;
