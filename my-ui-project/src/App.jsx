import { useState } from 'react';
import './App.css';

// Import all the components
import CapitalSearch from './CapitalSearch';
import CurrencySearch from './CurrencySearch';
import LanguageSearch from './LanguageSearch';
import CodeSearch from './CodeSearch';
import NameSearch from './NameSearch';

const SEARCH_OPTIONS = [
  { key: 'capital', label: 'Search by Capital', Component: CapitalSearch },
  { key: 'language', label: 'Search by Language', Component: LanguageSearch },
  { key: 'currency', label: 'Search by Currency', Component: CurrencySearch },
  { key: 'code', label: 'Search by Code', Component: CodeSearch },
  { key: 'name', label: 'Search by Name', Component: NameSearch },
];

function App() {
  const [searchType, setSearchType] = useState('capital');
  const SelectedComponent = SEARCH_OPTIONS.find(opt => opt.key === searchType)?.Component;

  return (
    <>
      <h1>Country Finder</h1>
      <p>Select a search method to find a country.</p>

      <select value={searchType} onChange={(e) => setSearchType(e.target.value)}>
        {SEARCH_OPTIONS.map(opt => (
          <option key={opt.key} value={opt.key}>
            {opt.label}
          </option>
        ))}
      </select>
      
      <hr />

      {SelectedComponent && <SelectedComponent />}

      {/* --- ADD THIS FOOTER HERE --- */}
      <footer style={{ marginTop: '50px', fontSize: '0.8rem', color: '#888' }}>
        <p>Powered by <strong>Go (Gin)</strong>, <strong>OpenAPI 3.0</strong>, and <strong>Docker</strong> on <strong>AWS App Runner</strong>.</p>
      </footer>
      {/* ---------------------------- */}
    </>
  );
}

export default App;