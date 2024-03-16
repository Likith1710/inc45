import { useRouter } from 'next/router';

export default function Results() {
  const router = useRouter();
  const { name, num1, num2 } = router.query;

  const addition = parseInt(num1) + parseInt(num2);
  const multiplication = parseInt(num1) * parseInt(num2);

  return (
    <div>
      <h1>Welcome back, {name}!</h1>
      <h2>Previous Search Results:</h2>
      <ul>
        <li>Addition: {addition}</li>
        <li>Multiplication: {multiplication}</li>
      </ul>
      <p>Enter new numbers:</p>
      <form>
        <label htmlFor="num1">Please enter first number:</label><br />
        <input type="number" id="num1" /><br />
        <label htmlFor="num2">Please enter second number:</label><br />
        <input type="number" id="num2" /><br />
        <button type="submit">Submit</button>
      </form>
    </div>
  );
}

