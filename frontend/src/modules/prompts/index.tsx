import React, { useState } from 'react';
import './index.css';
import { redirect, useNavigate } from 'react-router-dom';

const Prompts: React.FC = () => {
  const [age, setAge] = useState<number | ''>('');
  const [occupation, setOccupation] = useState<string>('');
  const [income, setIncome] = useState<number | ''>('');
  const [purpose, setPurpose] = useState<string[]>([]);
  const navigate = useNavigate();

  const handlePurposeChange = (value: string) => {
    setPurpose((prev) =>
      prev.includes(value) ? prev.filter((item) => item !== value) : [...prev, value]
    );
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    if (age === '' || occupation === '' || income === '' || purpose.length === 0) {
      alert('Please fill out all fields.');
      return;
    }

    console.log('Form Data:', { age, occupation, income, purpose });

    setAge('');
    setOccupation('');
    setIncome('');
    setPurpose([]);

    navigate('/plaid')
  };

  return (
    <div>
    <div className="dashboard-container">
      <h2 className="create-account-title">Tell Us More About Yourself</h2>
      <form className="dashboard-form" onSubmit={handleSubmit}>
        {/* Age */}
        <div className="form-group">
          <label htmlFor="age">1. What is your age?</label>
          <input
            type="number"
            id="age"
            value={age}
            onChange={(e) => setAge(e.target.value === '' ? '' : Number(e.target.value))}
            placeholder="Enter your age"
            required
          />
        </div>

        <div className="form-group">
          <label htmlFor="occupation">2. What is your occupation?</label>
          <select
            id="occupation"
            value={occupation}
            onChange={(e) => setOccupation(e.target.value)}
            required
          >
            <option value="" disabled>
              Select your occupation
            </option>
            <option value="Student">Student</option>
            <option value="Full-time">Full-time Employee</option>
            <option value="Part-time">Part-time Employee</option>
            <option value="None of the above">None of the above</option>
          </select>
        </div>

        <div className="form-group">
          <label htmlFor="income">3. What is your monthly income?</label>
          <select
            id="income"
            value={income}
            onChange={(e) => setIncome(e.target.value === '' ? '' : Number(e.target.value))}
            required
          >
            <option value="" disabled>
              Select your income
            </option>
            {[...Array(9)].map((_, i) => (
              <option key={i} value={(i + 1) * 10000}>
                ${((i + 1) * 10000).toLocaleString()}
              </option>
            ))}
            <option value='100000'>
                $100000+
            </option>
          </select>
        </div>

        <div className="form-group">
          <label>4. What are you looking for from this platform?</label>
          <div className="checkbox-group">
            {[
              { value: 'invest', label: 'Investing your money.' },
              { value: 'news', label: 'Keeping up with the latest financial news.' },
              { value: 'budget', label: 'Balancing your budget.' },
              { value: 'stock', label: 'Stock recommendations.' },
            ].map((option) => (
              <label key={option.value} className="checkbox-label">
                <input
                  type="checkbox"
                  value={option.value}
                  checked={purpose.includes(option.value)}
                  onChange={() => handlePurposeChange(option.value)}
                />
                <span className="checkbox-custom"></span>
                {option.label}
              </label>
            ))}
          </div>
        </div>

        <button type="submit" className="submit-button">
            Submit
        </button>
      </form>
    </div>
    <div>
    <button className="skip-button" onClick={() => navigate('/plaid')}>
    Skip
  </button>
    </div>
    </div>
  );
};

export default Prompts;
