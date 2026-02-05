import type { CurrenciesDetails } from './api-client';

interface CurrencyDisplayProps {
  currencies: { [key: string]: CurrenciesDetails };
}

function CurrencyDisplay({ currencies }: CurrencyDisplayProps) {
  return (
    <span>
      {Object.values(currencies).map(c => `${c.name} (${c.symbol})`).join(', ')}
    </span>
  );
}

export default CurrencyDisplay;