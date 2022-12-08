import { Paper, PaperProps } from "@mui/material";

export type SectionProps = PaperProps;

export const Section = (props: SectionProps) => {
  return (
    <Paper
      {...props}
      variant="outlined"
      sx={{
        padding: (theme) => theme.spacing(2),
        ...props.sx,
      }}
    />
  );
};
