import type { DemonymsDetails } from './api-client';

interface DeonymDetailsProps {
  deonyms: { [language: string]: DemonymsDetails };
}

function DeonymDetails({ deonyms }: DeonymDetailsProps) {
  return (
    <div style={{ padding: '1rem', fontFamily: 'sans-serif', color: '#eee' }}>
      <h3 style={{ marginBottom: '0.75rem', color: '#fff' }}>Demonyms</h3>
      <table style={{ width: '100%', borderCollapse: 'collapse' }}>
        <thead>
          <tr>
            <th style={thStyle}>Language</th>
            <th style={thStyle}>Female</th>
            <th style={thStyle}>Male</th>
          </tr>
        </thead>
        <tbody>
          {Object.entries(deonyms).map(([lang, { f, m }]) => (
            <tr key={lang}>
              <td style={tdStyle}>{lang.toUpperCase()}</td>
              <td style={tdStyle}>{f}</td>
              <td style={tdStyle}>{m}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

const thStyle: React.CSSProperties = {
  textAlign: 'left',
  padding: '8px 12px',
  backgroundColor: '#222',
  color: '#ccc',
  borderBottom: '1px solid #444',
};

const tdStyle: React.CSSProperties = {
  padding: '8px 12px',
  borderBottom: '1px solid #333',
};

export default DeonymDetails;
