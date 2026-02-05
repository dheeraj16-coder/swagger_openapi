

type GiniDisplayProps = {
  gini: { [year: string]: number };
};

function GiniDisplay({ gini }: GiniDisplayProps) {
  const getColor = (value: number) => {
    if (value < 30) return 'green';
    if (value < 40) return 'orange';
    return 'red';
  };

  return (
    <div>
      {Object.entries(gini).map(([year, value]) => (
        <div key={year}>
          <strong>{year}:</strong>{' '}
          <span style={{ color: getColor(value) }}>{value}</span>
        </div>
      ))}
    </div>
  );
}

export default GiniDisplay;
