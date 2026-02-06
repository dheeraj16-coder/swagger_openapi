import { useState, useEffect, useRef } from 'react';
import Header from './components/Header';
import { CountriesApi, Configuration } from './api-client';
import './App.css';

function App() {
  // Theme state
  const [theme, setTheme] = useState('light');
  
  // Search state
  const [searchType, setSearchType] = useState('capital');
  const [searchQuery, setSearchQuery] = useState('');
  const [selectedFields, setSelectedFields] = useState({
    name: true,
    capital: true,
    population: true,
    flags: true,
    region: false,
    currencies: false,
    languages: false,
    area: false
  });
  
  // Results state
  const [countries, setCountries] = useState([]);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState(null);

  // Tab indicator ref
  const tabSliderRef = useRef(null);
  const indicatorRef = useRef(null);

  // Apply theme to document
  useEffect(() => {
    document.documentElement.setAttribute('data-theme', theme);
  }, [theme]);

  // Update tab indicator position
  useEffect(() => {
    const updateIndicator = () => {
      const activeTab = tabSliderRef.current?.querySelector('.tab-option.active');
      const indicator = indicatorRef.current;
      
      if (activeTab && indicator && tabSliderRef.current) {
        const tabRect = activeTab.getBoundingClientRect();
        const sliderRect = tabSliderRef.current.getBoundingClientRect();
        
        indicator.style.width = `${tabRect.width}px`;
        indicator.style.height = `${tabRect.height}px`;
        indicator.style.left = `${tabRect.left - sliderRect.left}px`;
        indicator.style.top = `${tabRect.top - sliderRect.top}px`;
      }
    };

    updateIndicator();
    window.addEventListener('resize', updateIndicator);
    return () => window.removeEventListener('resize', updateIndicator);
  }, [searchType]);

  const toggleTheme = () => {
    setTheme(prev => prev === 'light' ? 'dark' : 'light');
  };

  const handleSearch = async () => {
    if (!searchQuery.trim()) {
      setError('Please enter a search term');
      return;
    }
    
    setIsLoading(true);
    setError(null);
    
    try {
      const apiClient = new CountriesApi(new Configuration({ basePath: '/v3.1' }));
      const fields = Object.keys(selectedFields).filter(key => selectedFields[key]);
      
      if (fields.length === 0) {
        setError('Please select at least one field');
        setIsLoading(false);
        return;
      }
      
      let response;
      switch(searchType) {
        case 'capital':
          response = await apiClient.getCountryByCapital(searchQuery, fields);
          break;
        case 'name':
          response = await apiClient.getCountryByName(searchQuery, fields);
          break;
        case 'currency':
          response = await apiClient.getCountryByCurrency(searchQuery, fields);
          break;
        case 'language':
          response = await apiClient.getCountryByLanguage(searchQuery, fields);
          break;
        case 'code':
          response = await apiClient.getCountryByCode(searchQuery, fields);
          break;
        default:
          response = await apiClient.getCountryByCapital(searchQuery, fields);
      }
      
      setCountries(response.data || []);
      if (!response.data || response.data.length === 0) {
        setError('No countries found');
      }
    } catch (err) {
      console.error('Search error:', err);
      setError('Could not find countries. Please try again.');
      setCountries([]);
    } finally {
      setIsLoading(false);
    }
  };

  const handleFieldToggle = (field) => {
    setSelectedFields(prev => ({
      ...prev,
      [field]: !prev[field]
    }));
  };

  const handleTabClick = (type) => {
    setSearchType(type);
    setSearchQuery('');
    setCountries([]);
    setError(null);
  };

  return (
    <div className="container">
      <Header theme={theme} toggleTheme={toggleTheme} />
      
      {/* Search Card */}
      <div className="search-card">
        {/* Tab Navigation */}
        <div className="tab-slider-container">
          <div className="tab-slider" ref={tabSliderRef}>
            <div className="tab-indicator" ref={indicatorRef}></div>
            <button 
              className={`tab-option ${searchType === 'capital' ? 'active' : ''}`}
              onClick={() => handleTabClick('capital')}
            >
              <span>🏛️</span>Capital
            </button>
            <button 
              className={`tab-option ${searchType === 'name' ? 'active' : ''}`}
              onClick={() => handleTabClick('name')}
            >
              <span>🔤</span>Name
            </button>
            <button 
              className={`tab-option ${searchType === 'language' ? 'active' : ''}`}
              onClick={() => handleTabClick('language')}
            >
              <span>💬</span>Language
            </button>
            <button 
              className={`tab-option ${searchType === 'currency' ? 'active' : ''}`}
              onClick={() => handleTabClick('currency')}
            >
              <span>💰</span>Currency
            </button>
            <button 
              className={`tab-option ${searchType === 'code' ? 'active' : ''}`}
              onClick={() => handleTabClick('code')}
            >
              <span>🔢</span>Code
            </button>
          </div>
        </div>

        {/* Search Input */}
        <div className="input-group">
          <span className="input-icon">🔍</span>
          <input 
            type="text" 
            className="premium-input"
            value={searchQuery}
            onChange={(e) => setSearchQuery(e.target.value)}
            onKeyPress={(e) => e.key === 'Enter' && handleSearch()}
            placeholder={`Enter ${searchType} (e.g., ${
              searchType === 'capital' ? 'London, Paris, Tokyo' :
              searchType === 'name' ? 'United States, France' :
              searchType === 'language' ? 'English, Spanish' :
              searchType === 'currency' ? 'USD, EUR, GBP' :
              'US, FR, JP'
            })`}
          />
        </div>

        {/* Field Selectors */}
        <div className="field-selectors">
          <label className="field-container">
            <input 
              type="checkbox" 
              checked={selectedFields.name}
              onChange={() => handleFieldToggle('name')}
            />
            <div className="switch">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2">
                <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"></path>
                <circle cx="12" cy="7" r="4"></circle>
              </svg>
            </div>
            <span className="field-label">Name</span>
          </label>

          <label className="field-container">
            <input 
              type="checkbox" 
              checked={selectedFields.capital}
              onChange={() => handleFieldToggle('capital')}
            />
            <div className="switch">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2">
                <path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"></path>
              </svg>
            </div>
            <span className="field-label">Capital</span>
          </label>

          <label className="field-container">
            <input 
              type="checkbox" 
              checked={selectedFields.population}
              onChange={() => handleFieldToggle('population')}
            />
            <div className="switch">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2">
                <path d="M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2"></path>
                <circle cx="9" cy="7" r="4"></circle>
                <path d="M23 21v-2a4 4 0 0 0-3-3.87"></path>
                <path d="M16 3.13a4 4 0 0 1 0 7.75"></path>
              </svg>
            </div>
            <span className="field-label">Population</span>
          </label>

          <label className="field-container">
            <input 
              type="checkbox" 
              checked={selectedFields.flags}
              onChange={() => handleFieldToggle('flags')}
            />
            <div className="switch">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2">
                <path d="M4 15s1-1 4-1 5 2 8 2 4-1 4-1V3s-1 1-4 1-5-2-8-2-4 1-4 1z"></path>
                <line x1="4" y1="22" x2="4" y2="15"></line>
              </svg>
            </div>
            <span className="field-label">Flags</span>
          </label>

          <label className="field-container">
            <input 
              type="checkbox" 
              checked={selectedFields.region}
              onChange={() => handleFieldToggle('region')}
            />
            <div className="switch">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2">
                <circle cx="12" cy="12" r="10"></circle>
                <line x1="2" y1="12" x2="22" y2="12"></line>
                <path d="M12 2a15.3 15.3 0 0 1 4 10 15.3 15.3 0 0 1-4 10 15.3 15.3 0 0 1-4-10 15.3 15.3 0 0 1 4-10z"></path>
              </svg>
            </div>
            <span className="field-label">Region</span>
          </label>

          <label className="field-container">
            <input 
              type="checkbox" 
              checked={selectedFields.currencies}
              onChange={() => handleFieldToggle('currencies')}
            />
            <div className="switch">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2">
                <line x1="12" y1="1" x2="12" y2="23"></line>
                <path d="M17 5H9.5a3.5 3.5 0 0 0 0 7h5a3.5 3.5 0 0 1 0 7H6"></path>
              </svg>
            </div>
            <span className="field-label">Currency</span>
          </label>

          <label className="field-container">
            <input 
              type="checkbox" 
              checked={selectedFields.languages}
              onChange={() => handleFieldToggle('languages')}
            />
            <div className="switch">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2">
                <path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"></path>
              </svg>
            </div>
            <span className="field-label">Languages</span>
          </label>

          <label className="field-container">
            <input 
              type="checkbox" 
              checked={selectedFields.area}
              onChange={() => handleFieldToggle('area')}
            />
            <div className="switch">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2">
                <rect x="3" y="3" width="18" height="18" rx="2" ry="2"></rect>
              </svg>
            </div>
            <span className="field-label">Area</span>
          </label>
        </div>

        {/* Search Button */}
        <div className="search-btn-wrapper">
          <button className={`button-3d ${isLoading ? 'loading' : ''}`} onClick={handleSearch}>
            <span className="shadow"></span>
            <span className="edge"></span>
            <span className="front">
              <span className="button-text">🚀 Search Countries</span>
              <span className="spinner"></span>
            </span>
          </button>
        </div>
      </div>

      {/* Results */}
      <div className="results-header">
        <div className="results-count">
          {error ? <span style={{color: '#ff3b30'}}>{error}</span> : 
           isLoading ? 'Searching...' :
           countries.length > 0 ? `Found ${countries.length} ${countries.length === 1 ? 'country' : 'countries'}` :
           'Ready to search'}
        </div>
      </div>

      <div className="countries-grid">
        {isLoading ? (
          // Loading skeletons
          Array(3).fill(0).map((_, i) => (
            <div key={i} className="skeleton-card">
              <div className="skeleton-flag"></div>
              <div className="skeleton-content">
                <div className="skeleton-line"></div>
                <div className="skeleton-line"></div>
                <div className="skeleton-line"></div>
                <div className="skeleton-line"></div>
              </div>
            </div>
          ))
        ) : (
          countries.map(country => (
            <div key={country.cca3 || Math.random()} className="country-card">
              {country.flags && (
                <div className="flag-container">
                  <img src={country.flags.png} alt={country.name?.common} className="country-flag" />
                </div>
              )}
              <div className="country-info">
                <div className="country-header">
                  <div className="country-name">{country.name?.common || 'Unknown'}</div>
                  {country.region && (
                    <div className="country-region">
                      {country.region} {country.subregion && `· ${country.subregion}`}
                      {country.independent && <span className="badge">Independent</span>}
                    </div>
                  )}
                </div>
                <div className="country-stats">
                  {country.capital && (
                    <div className="stat-item">
                      <span className="stat-label">Capital</span>
                      <span className="stat-value">{Array.isArray(country.capital) ? country.capital[0] : country.capital}</span>
                    </div>
                  )}
                  {country.population && (
                    <div className="stat-item">
                      <span className="stat-label">Population</span>
                      <span className="stat-value">{country.population.toLocaleString()}</span>
                    </div>
                  )}
                  {country.currencies && (
                    <div className="stat-item">
                      <span className="stat-label">Currency</span>
                      <span className="stat-value">
                        {Object.values(country.currencies).map(c => `${c.symbol || ''} ${c.name}`).join(', ')}
                      </span>
                    </div>
                  )}
                  {country.languages && (
                    <div className="stat-item">
                      <span className="stat-label">Languages</span>
                      <span className="stat-value">
                        {Object.values(country.languages).join(', ')}
                      </span>
                    </div>
                  )}
                  {country.area && (
                    <div className="stat-item">
                      <span className="stat-label">Area</span>
                      <span className="stat-value">{country.area.toLocaleString()} km²</span>
                    </div>
                  )}
                </div>
              </div>
            </div>
          ))
        )}
      </div>
    </div>
  );
}

export default App;
