import { Typography, TypographyProps } from "@mui/material";

export type LabelProps = TypographyProps;

export const Label = (props: TypographyProps) => {
  return <Typography variant="h6" component="span" {...props} />;
};
