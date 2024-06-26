import { useRef } from 'react';

function NewC() {
  const inputRef = useRef(null);

  const handleClick = () => {
    inputRef.current.value = 'New value';
  };

  return (
    <div>
      <input ref={inputRef} type="text" />
      <button onClick={handleClick}>Change input value</button>
    </div>
  );
}
export default NewC;