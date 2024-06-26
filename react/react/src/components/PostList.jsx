import React from "react";
import PostItem from "./PostItem";

const PostList = ({posts, title}) => {
  return(
  <>
  <h1 style={{textAlign:"center"}}>
    {title}
  </h1>
  {posts.map((item,index)=>
     <PostItem number={index+1} post={item} key={item.id}/>)}
  </>
  )
}

export default PostList;