import React, { useState } from 'react';
import { CountriesApi, Configuration } from './api-client';
import type { AllresponsebodyInner as Country } from './api-client';
import CountryDetails from './DisplayDetails';

const availableFields = ['name', 'capital', 'population', 'flags', 'currencies','tld','cca2','cca3','region','subregion','languages','timezones','borders','independent','unMember','cioc','continents', 'area', 'landlocked', 'gini', 'demonyms', 'translations', 'maps', 'car', 'status', 'altSpellings', 'idd', 'capitalInfo', 'coatOfArms'];

function CodeSearch() {
  const [code, setCode] = useState('');
  const [countries, setCountries] = useState<Country[]>([]);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  
  // We'll use an object to keep track of which fields are checked.
  const [selectedFields, setSelectedFields] = useState({
    name: true,
    capital: true,
    population: true,
    flags: true,
    currencies: false,
    tld: false,
    cca2: false,
    cca3: false,
    region: false,
    subregion: false,
    languages: false,
    timezones: false,
    borders: false,
    independent: false,
    unMember: false,
    cioc: false,
    continents: false,  
    area: false,
    landlocked: false,
    gini: false,
    demonyms: false,
    translations: false,
    maps: false,
    car: false,
    altSpellings: false,
    idd: false,
    capitalInfo: false,
    coatOfArms: false,
  });
  const apiClient = new CountriesApi(new Configuration({ basePath: 'http://localhost:8080/v3.1' }));

  const handleCheckboxChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const { name, checked } = event.target;
    setSelectedFields(prevFields => ({
      ...prevFields,
      [name]: checked,
    }));
  };

  const handleSearch = () => {
    setIsLoading(true);
    setError(null);
    setCountries([]);


    // Get an array of the keys from our state object where the value is true
    const fieldsToRequest = Object.keys(selectedFields).filter(key => selectedFields[key]);
    
    // Check if any fields are selected to avoid sending an empty parameter
    if (fieldsToRequest.length === 0) {
        setError("Please select at least one field to display.");
        setIsLoading(false);
        return;
    }

    apiClient.getCountryByCode(code, fieldsToRequest) // Pass the selected fields
      .then(response => {
        setCountries(response.data);
      })
      .catch(err => {
        setError('Could not find a country with that code.');
        console.error('API Error:', err);
      })
      .finally(() => {
        setIsLoading(false);
      });
  };

  return (
    <div>
      <h2>Search by Capital City</h2>
      <input
        type="text"
        value={code}
        onChange={(e) => setCode(e.target.value)}
        placeholder="e.g., WAR"
      />
      <button onClick={handleSearch} disabled={isLoading}>
        {isLoading ? 'Searching...' : 'Search'}
      </button>

      {/* --- NEW FEATURE 1: Checkboxes for field selection --- */}
      <div style={{ margin: '10px 0' }}>
        <strong>Fields to include:</strong>
        {availableFields.map(field => (
          <label key={field} style={{ margin: '0 10px' }}>
            <input
              type="checkbox"
              name={field}
              checked={selectedFields[field]}
              onChange={handleCheckboxChange}
            />
            {field}
          </label>
        ))}
      </div>

      {isLoading && <p>Loading...</p>}
      {error && <p style={{ color: 'red' }}>{error}</p>}

       {countries.length > 0 && (
        <div>
          <h3>Results:</h3>
          {countries.map(country => (
            // For each country, render our new dynamic details component
            <CountryDetails key={country.cca3} country={country} />
          ))}
        </div>
      )}
    </div>
  );
}


export default CodeSearch;