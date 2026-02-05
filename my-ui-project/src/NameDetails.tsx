import type { NameDetails} from './api-client';

interface  NameDisplayProps{
  names: NameDetails;
};

function NameDisplay({ names }: NameDisplayProps) {
  // Guard against the case where names is null or undefined
  if (!names) {
    return null;
  }

  return (
    <div>
      <p>
        <strong>Official:</strong> {names.official || 'N/A'}
      </p>
      <p>
        <strong>Common:</strong> {names.common || 'N/A'}
      </p>

      {names.nativeName && (
        <div>
          <strong>Native Names:</strong>
          <ul>
            {Object.entries(names.nativeName).map(([langCode, nameDetail]) => (
              <li key={langCode}>
                {langCode}: {nameDetail.common} ({nameDetail.official})
              </li>
            ))}
          </ul>
        </div>
      )}
    </div>
  );
}

export default NameDisplay;