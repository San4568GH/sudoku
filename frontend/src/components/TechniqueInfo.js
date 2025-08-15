import React from 'react';
import { Card, Title, Text, Button, Stack } from '@mantine/core';

const definitions = {
  "Naked Single": "A Naked Single is a cell that has only one possible candidate. This is the easiest and most common Sudoku technique.",
  "Hidden Single": "A Hidden Single is a cell where a certain number can only be placed in one cell within a row, column, or 3x3 box, although the cell itself has other candidates.",
  "Hidden Single (Simplified)": "A Hidden Single is a cell where a certain number can only be placed in one cell within a row, column, or 3x3 box, although the cell itself has other candidates."
};

function TechniqueInfo({ technique, onClose }) {
  if (!technique) return null;

  return (
    <Card
      style={{
        background: 'rgba(255, 255, 255, 0.1)',
        backdropFilter: 'blur(10px)',
        border: '1px solid rgba(255, 255, 255, 0.2)',
        color: 'white',
        height: '100%',
        display: 'flex',
        flexDirection: 'column'
      }}
      p="xl"
      radius="md"
    >
      <Stack justify="space-between" style={{ flex: 1 }}>
        <div>
          <Title order={3} mb="md">{technique}</Title>
          <Text size="md" style={{ lineHeight: 1.6 }}>
            {definitions[technique] || "No definition available."}
          </Text>
        </div>
        <Button onClick={onClose} variant="outline" color="gray" mt="md">
          Close
        </Button>
      </Stack>
    </Card>
  );
}

export default TechniqueInfo;
