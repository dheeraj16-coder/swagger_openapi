// In src/CountryDetails.tsx

import type { AllresponsebodyInner as Country } from './api-client/index.ts';

// Import our new specialist components

import CurrencyDisplay from './CurrencyDisplay.tsx';
import MapLinks from './MapLinks.tsx';
import GiniDetails from './GiniDetails.tsx';
import CapitalInfo from './CapitalInfo.tsx';
import IddDetailsDisplay from './IddDetails.tsx';
import CarDetailsDisplay from './CarDetails.tsx';
import TranslationDisplay from './TranslationDetails.tsx';
import DemonymDetails from './DeonymsDetails.tsx';
import LanguageDetails from './LanguageDetails.tsx';  


function CountryDetails({ country }: { country: Country }) {
  return (
    <div style={{ textAlign: 'left', border: '1px solid grey', padding: '10px', marginTop: '10px' }}>
      {country.name && <h3>{country.name.common}</h3>}

      {Object.entries(country).map(([key, value]) => {
      if (value === null || value === undefined || ['name'].includes(key)) {
               return null;
  }

        let content;
        
        // This switch statement decides how to render each value based on its key.
        console.log(key, value, typeof value)
        switch (key) {
          case 'flags':
          case 'coatOfArms':
            content = <img src={value.png} alt={value.alt || key} width="100" />;
            break;
          case 'currencies':
            content = <CurrencyDisplay currencies={value} />;
            break;
          case 'gini':
            content = <GiniDetails gini={value} />;
            break;
          case 'capitalInfo':
           content = <CapitalInfo capitalInfo={value} />;
           break;
          case 'maps':
            content = <MapLinks maps={value} />;
            break;
          case 'idd':
            content = <IddDetailsDisplay idd={value} />;
            break;
          case 'car':
            content = <CarDetailsDisplay car={value} />;
            break;
          case 'translations':
            content = <TranslationDisplay translations={value} />;
            break;
          case 'demonyms':
            content = <DemonymDetails deonyms={value} />;
            break;
          case 'languages':
            content = <LanguageDetails languages={value} />;
            break;
          default:
             if (typeof value === 'boolean') {
      content = <span>{value ? '✅' : '❌'}</span>;
    } else {
      content = <span>{Array.isArray(value) ? value.join(', ') : String(value)}</span>;
    }
        }
        
        return (
          <div key={key}>
            <strong>{key}:</strong> {content}
          </div>
        );
      })}
    </div>
  );
}

export default CountryDetails;