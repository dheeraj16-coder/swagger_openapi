import type { NativeNameDetail } from './api-client';

interface TranslationDisplayProps {
  translations: { [key: string]: NativeNameDetail };
}

function TranslationDisplay({ translations }: TranslationDisplayProps) {
  const entries = Object.entries(translations);

  return (
    <div style={{ display: 'flex', flexWrap: 'wrap', gap: '0.5rem' }}>
      {entries.map(([lang, t]) => (
        <div
          key={lang}
          style={{
            border: '1px solid #ccc',
            borderRadius: '6px',
            padding: '0.3rem 0.6rem',
            backgroundColor: 'black',
            fontSize: '0.9rem',
            whiteSpace: 'nowrap',
          }}
          title={`Language code: ${lang}`}
        >
          <strong>{t.official}</strong> <small>({t.common})</small>
        </div>
      ))}
    </div>
  );
}

export default TranslationDisplay;
