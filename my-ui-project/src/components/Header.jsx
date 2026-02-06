import React from 'react';
import './Header.css';

const Header = ({ theme, toggleTheme }) => {
  return (
    <header className="header">
      <div className="header-content">
        <div className="logo-section">
          {/* Theme Toggle as Logo */}
          <label className="theme-toggle-logo themeToggle st-sunMoonThemeToggleBtn">
            <input 
              type="checkbox" 
              className="themeToggleInput" 
              checked={theme === 'dark'}
              onChange={toggleTheme}
            />
            <svg width="36" height="36" viewBox="0 0 20 20" fill="currentColor" stroke="none">
              <mask id="moon-mask">
                <rect x="0" y="0" width="20" height="20" fill="white"></rect>
                <circle cx="11" cy="3" r="8" fill="black"></circle>
              </mask>
              <circle className="sunMoon" cx="10" cy="10" r="8" mask="url(#moon-mask)"></circle>
              <g>
                <circle className="sunRay sunRay1" cx="18" cy="10" r="1.5"></circle>
                <circle className="sunRay sunRay2" cx="14" cy="16.928" r="1.5"></circle>
                <circle className="sunRay sunRay3" cx="6" cy="16.928" r="1.5"></circle>
                <circle className="sunRay sunRay4" cx="2" cy="10" r="1.5"></circle>
                <circle className="sunRay sunRay5" cx="6" cy="3.072" r="1.5"></circle>
                <circle className="sunRay sunRay6" cx="14" cy="3.072" r="1.5"></circle>
              </g>
            </svg>
          </label>
          <div className="logo-text">
            <h1>Country Explorer</h1>
            <p>Discover nations across the globe</p>
          </div>
        </div>
      </div>
    </header>
  );
};

export default Header;
