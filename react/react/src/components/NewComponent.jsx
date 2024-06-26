
import {useState , useEffect} from 'react';

const  NewComponent= function(){
let [statement,setStatement] = useState(0);
useEffect(useEffectFunction,[statement])
 function hey(){
  setStatement(statement+1)
}
 function useEffectFunction(){
  setTimeout(()=> alert('Holla!'),2000)
 }
return(
      <>
          <button onClick={hey}>click</button>
          <button>{statement}</button>
          <p></p>
      </>

  );
}

export default NewComponent;