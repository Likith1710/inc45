import { useState } from 'react';
import { useRouter } from 'next/router';

export default function Home() {
  const [name, setName] = useState('');
  const [num1, setNum1] = useState('');
  const [num2, setNum2] = useState('');
  const router = useRouter();

  const handleSubmit = (e) => {
    e.preventDefault();
    router.push({
      pathname: '/results',
      query: { name, num1, num2 },
    });
  };

  return (
    <div>
      <h1>Welcome to the Calculator App!</h1>
      <form onSubmit={handleSubmit}>
        <label htmlFor="name">Please enter your name:</label><br />
        <input type="text" id="name" value={name} onChange={(e) => setName(e.target.value)} /><br />
        <label htmlFor="num1">Please enter first number:</label><br />
        <input type="number" id="num1" value={num1} onChange={(e) => setNum1(e.target.value)} /><br />
        <label htmlFor="num2">Please enter second number:</label><br />
        <input type="number" id="num2" value={num2} onChange={(e) => setNum2(e.target.value)} /><br />
        <button type="submit">Submit</button>
      </form>
    </div>
  );
}

