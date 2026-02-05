import type { IddDetails } from './api-client';

interface IddDetailsProps {
  idd: IddDetails;
}

function IddDetailsDisplay({ idd }: IddDetailsProps) {
  const combinedIds = idd.suffixes ? idd.suffixes.map(suffix => idd.root + suffix) : [];
  return (
  <div>
    {combinedIds.map((id, index) => (
      <div key={index}>{id}</div>
    ))}
  </div>
);
}
export default IddDetailsDisplay;