import { useState } from 'react';
import './App.css';

// Import all the components you want to be able to switch between
import CapitalSearch from './CapitalSearch';
import CurrencySearch from './CurrencySearch';
import LanguageSearch from './LanguageSearch';
import CodeSearch from './CodeSearch';
import NameSearch from './NameSearch';

// Step 1: DEFINE your search options in a configuration array.
const SEARCH_OPTIONS = [
  { key: 'capital', label: 'Search by Capital', Component: CapitalSearch },
  { key: 'language', label: 'Search by Language', Component: LanguageSearch },
  { key: 'currency', label: 'Search by Currency', Component: CurrencySearch },
  { key: 'code', label: 'Search by Code', Component: CodeSearch },
  { key: 'name', label: 'Search by Name', Component: NameSearch },
];

function App() {
  // Step 2: TRACK the current selection in state.
  // 'capital' is the default selected option.
  const [searchType, setSearchType] = useState('capital');

  // Find the component that matches the currently selected searchType.
  const SelectedComponent = SEARCH_OPTIONS.find(opt => opt.key === searchType)?.Component;

  return (
    <>
      <h1>Country Finder</h1>
      <p>Select a search method to find a country.</p>

      <select value={searchType} onChange={(e) => setSearchType(e.target.value)}>
        {/* The dropdown options are now generated from our config array */}
        {SEARCH_OPTIONS.map(opt => (
          <option key={opt.key} value={opt.key}>
            {opt.label}
          </option>
        ))}
      </select>
      
      <hr />

      {/* Step 3: RENDER the selected component dynamically. */}
      {SelectedComponent && <SelectedComponent />}
    </>
  );
}

export default App;