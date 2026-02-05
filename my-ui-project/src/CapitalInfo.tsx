import type { AllresponsebodyInnerCapitalInfo } from './api-client';

interface CapitalInfoProps {
 capitalInfo:  AllresponsebodyInnerCapitalInfo | undefined;
}

function CapitalInfo({ capitalInfo }: CapitalInfoProps) {
  if (!capitalInfo) return null;
    return (
        <span>
        {capitalInfo.latlng ? `Latitude: ${capitalInfo.latlng[0]}, Longitude: ${capitalInfo.latlng[1]}` : 'No capital info available'}
        </span>
    );
}
export default CapitalInfo;